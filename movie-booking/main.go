package main

import (
	"movie-booking/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	routes.MovieRoute(r)

	http.ListenAndServe(":8080", r)
}
