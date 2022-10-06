package main

import (
	"log"
	"mock/application"
	"net/http"

	"github.com/rs/cors"
)
var corsHandler = cors.New(cors.Options{
	AllowedOrigins:   []string{"http://localhost:3006"},
	AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
	AllowedHeaders:   []string{"Origin", "Accept", "Content-Type", "X-Requested-With"},
	AllowCredentials: true,
	MaxAge:           0,
	Debug:            true,
})

var handler = corsHandler.Handler(application.SetupRouter())

var port = "4000"

func main () {
	err := http.ListenAndServe(":"+port, handler)
	if err != nil {
		log.Fatal(err)
	}
}