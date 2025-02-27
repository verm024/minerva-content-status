package main

import (
	"log"
	"minerva-content-status/db"
	"minerva-content-status/router"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	envLoadError := godotenv.Load()
	if envLoadError != nil {
		log.Fatalf("Error loading .env file: %v", envLoadError)
	}

	db, _ := db.Connect()

	e := echo.New()
	router.Initialize(e, db)
	e.Logger.Fatal(e.Start(":5020"))
}
