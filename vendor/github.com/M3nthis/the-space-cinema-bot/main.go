package main

import (
	"log"

	"github.com/M3nthis/the-space-cinema-bot/app"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	app.StartApp()
}
