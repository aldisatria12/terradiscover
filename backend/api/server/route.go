package server

import (
	"database/sql"
	"net/http"

	"github.com/aldisatria12/terradiscover/handler"
	"github.com/aldisatria12/terradiscover/middleware"
	"github.com/aldisatria12/terradiscover/repository"
	"github.com/aldisatria12/terradiscover/service"
	"github.com/aldisatria12/terradiscover/util/logger"
	"github.com/gin-gonic/gin"
)

type Route struct {
	userHandler    *handler.UserHandler
	contactHandler *handler.ContactHandler
}

func NewRoute(db *sql.DB) Route {
	dataStore := repository.NewDataStore(db)
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(dataStore, userRepository)
	userHandler := handler.NewUserHandler(userService)

	contactRepository := repository.NewContactRepository(db)
	contactService := service.NewContactService(dataStore, contactRepository)
	contactHandler := handler.NewContactHandler(contactService)

	return Route{
		userHandler:    &userHandler,
		contactHandler: &contactHandler,
	}
}

func (route *Route) SetRoutes() http.Handler {
	logrusLogger := logger.NewLogger()
	logger.SetLogger(logrusLogger)
	r := gin.New()
	r.ContextWithFallback = true
	r.Use(middleware.ErrorMiddleware, middleware.LoggerMiddleware(), gin.Recovery(),
		middleware.CORS())

	r.POST("/auth/login", route.userHandler.Login)
	r.POST("/auth/register", route.userHandler.Register)
	r.GET("/contact", middleware.AuthMiddleware(), route.contactHandler.GetContact)
	r.GET("/contact/:id", middleware.AuthMiddleware(), route.contactHandler.GetContactById)
	r.POST("/contact/insert", middleware.AuthMiddleware(), route.contactHandler.InsertContact)
	r.PUT("/contact/edit", middleware.AuthMiddleware(), route.contactHandler.EditContact)

	return r
}
