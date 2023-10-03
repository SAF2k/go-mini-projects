package main

import (
	"go-bookstore/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()

	// Define your CORS options
	c := cors.AllowAll()

	// Wrap your router with CORS middleware
	handler := c.Handler(r)

	// Register your routes
	routes.RegisterBookStoreRoutes(r)

	// Start the server with the CORS-wrapped handler
	log.Fatal(http.ListenAndServe(":8080", handler))
}
