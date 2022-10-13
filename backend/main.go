package main

import (
	"log"
	"mock/application"
	"net/http"
)


var port = "5000"

func main () {
	err := http.ListenAndServe(":"+port, application.SetupRouter())
	if err != nil {
		log.Fatal(err)
	}
}