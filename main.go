package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/s19835/auth-app-go/handlers"
)

func main() {
	r := mux.NewRouter()

	// public routes
	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	log.Println("Server started on :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
