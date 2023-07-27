package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/k56-Live/my-web-app/internal/handlers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/login", handlers.LoginHandler)
	r.HandleFunc("/dashboard", handlers.AuthMiddleware(handlers.DashboardHandler))

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
