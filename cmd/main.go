package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to My Web App!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the About page.")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contact us at contact@example.com")
}

func main() {
	// Create a new router using Gorilla Mux
	r := mux.NewRouter()

	// Define the routes and their corresponding handler functions
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/about", aboutHandler)
	r.HandleFunc("/contact", contactHandler)

	// Create a new HTTP server and bind the router to it
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Start the server
	fmt.Println("Server listening on port 8080...")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
