package main

import (
	"Search/internal/server"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed reading env file")
	}

	server.Start()
}