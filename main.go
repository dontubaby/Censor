package main

import (
	"Skillfactory/Censor/internal/api"
	middleware "github.com/dontubaby/mware"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	// Инициализация API
	api := api.New()
	//Инициализация роутера и подключение к нему middleware
	router := api.Router()
	router.Use(middleware.RequestIDMiddleware, middleware.LoggingMiddleware)

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("cant loading .env file")
	}
	port := os.Getenv("PORT")

	log.Println("Server gonews APP start working at port :4000!")
	err = http.ListenAndServe(port, router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
