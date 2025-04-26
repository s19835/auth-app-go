package main

import (
	"github.com/gorilla/mux"
	"github.com/s19835/auth-app-go/handlers"
)

func main() {
	r := mux.NewRouter()

	// public routes
	r.HandleFunc("/register", handlers.Register).Methods("POST")
}
