package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// InitRoutes initializes the routes
func InitRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", Hello).Methods("GET")
	return r
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}
