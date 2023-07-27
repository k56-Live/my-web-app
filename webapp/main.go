package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/username/my-web-app/webapp/internal/handlers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/dashboard", handlers.DashboardHandler)

	fmt.Println("Web application service is running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
