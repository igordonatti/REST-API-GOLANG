package main

import (
	"net/http"

	"github.com/igordonatti/REST-API/controllers"
	"github.com/igordonatti/REST-API/models"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	handler := controllers.New()

	server := &http.Server{
		Addr:    "0.0.0.0:8008",
		Handler: handler,
	}

	models.ConnectToDatabase()

	server.ListenAndServe()
}
