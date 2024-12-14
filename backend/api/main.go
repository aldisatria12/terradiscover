package main

import (
	"log"

	"github.com/aldisatria12/terradiscover/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf(err.Error())
	}
	server.Start()
}
