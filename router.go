package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vishals11/gouser/controller"
)

// InitRoutes initializes the routes
func InitRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", Hello).Methods("GET")
	r.HandleFunc("/user/signup", controller.SignUp).Methods("POST")
	return r
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}
