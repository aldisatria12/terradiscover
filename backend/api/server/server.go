package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aldisatria12/terradiscover/database"
)

func Start() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Fail to connect DB: %v", err.Error())
	}
	defer db.Close()

	route := NewRoute(db)

	handler := route.SetRoutes()

	addr := os.Getenv("BASE_URL")

	srv := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	log.Printf("Server running on %s", addr)

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("error server listen and serve: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	log.Println("shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	<-ctx.Done()
	log.Println("timeout of 5 seconds")

	log.Println("server exiting")

}
