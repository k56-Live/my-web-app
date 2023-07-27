package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/username/my-web-app/authentication/internal/handlers"
)

func main() {
	r := mux.NewRouter()
	authHandler := handlers.NewAuthHandler()

	r.HandleFunc("/login", authHandler.Login).Methods("POST")
	r.HandleFunc("/logout", authHandler.Logout).Methods("POST")

	server := &http.Server{
		Addr:    ":8081",
		Handler: r,
	}

	fmt.Println("Authentication service is running on http://localhost:8081")
	log.Fatal(server.ListenAndServe())
}
