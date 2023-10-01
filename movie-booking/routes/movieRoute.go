package routes

import (
	"movie-booking/controllers"

	"github.com/gorilla/mux"
)

func MovieRoute(r *mux.Router) {
	// Movie routes
	r.HandleFunc("/movies", controllers.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id:[0-9]+}", controllers.GetMovie).Methods("GET")
	r.HandleFunc("/movies", controllers.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id:[0-9]+}", controllers.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id:[0-9]+}", controllers.DeleteMovie).Methods("DELETE")
}
