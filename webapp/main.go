package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/k56-Live/my-web-app/webapp/internal/handlers"
)

func main() {
	r := mux.NewRouter()
	homeHandler := handlers.NewHomeHandler()

	r.HandleFunc("/", homeHandler.Index).Methods("GET")
	r.HandleFunc("/dashboard", homeHandler.Dashboard).Methods("GET")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	fmt.Println("Web application is running on http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}
