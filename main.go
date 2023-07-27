package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/username/my-web-app/authentication/internal/handlers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/login", handlers.AuthHandler)

	fmt.Println("Authentication service is running on http://localhost:8081")
	http.ListenAndServe(":8081", r)
}
