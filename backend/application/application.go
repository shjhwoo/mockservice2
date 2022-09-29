package application

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewHttpHandler() http.Handler {
	mux := mux.NewRouter()
	mux.HandleFunc("/", healthCheck).Methods("GET")
	mux.HandleFunc("/sso/login", ssologinHandler).Methods("GET")
	mux.HandleFunc("/callback", callbackHandler)
	mux.HandleFunc("/admin", getadminservice).Methods("GET")
	mux.HandleFunc("/all", getnormalservice).Methods("GET")
	return mux
}


