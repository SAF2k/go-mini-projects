// movie_controller.go

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Year  int    `json:"year"`
}

var movies []Movie

// GetMovies retrieves all movies.
func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// GetMovie retrieves a single movie by ID.
func GetMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Find the movie with the given ID and return it
	for _, movie := range movies {
		if fmt.Sprintf("%d", movie.ID) == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

	// If movie not found, return a 404 error
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Movie not found")
}

// CreateMovie creates a new movie.
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	var newMovie Movie
	if err := json.NewDecoder(r.Body).Decode(&newMovie); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Assign a unique ID and add the new movie to the slice
	newMovie.ID = len(movies) + 1
	movies = append(movies, newMovie)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newMovie)
}

// UpdateMovie updates an existing movie by ID.
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Find the movie with the given ID
	for i, movie := range movies {
		if fmt.Sprintf("%d", movie.ID) == id {
			var updatedMovie Movie
			if err := json.NewDecoder(r.Body).Decode(&updatedMovie); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// Update the movie's information
			updatedMovie.ID = movie.ID
			movies[i] = updatedMovie

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedMovie)
			return
		}
	}

	// If movie not found, return a 404 error
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Movie not found")
}

// DeleteMovie deletes a movie by ID.
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Find the movie with the given ID and delete it
	for i, movie := range movies {
		if fmt.Sprintf("%d", movie.ID) == id {
			movies = append(movies[:i], movies[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	// If movie not found, return a 404 error
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Movie not found")
}
