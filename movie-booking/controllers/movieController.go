package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Director represents the director of a movie.
type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// Movie represents a movie with an ID, title, year, and director.
type Movie struct {
	ID       int      `json:"id"`
	Title    string   `json:"title"`
	Year     int      `json:"year"`
	Director Director `json:"director"`
}

var movies []Movie // Slice for storing movies

// respondWithError writes an error response with the given status code and message.
func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	fmt.Fprintf(w, message)
}

// writeJSONResponse writes a JSON response with the given data and status code.
func writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// GetMovies retrieves all movies.
func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	writeJSONResponse(w, http.StatusOK, movies)
}

// GetMovie retrieves a single movie by ID.
func GetMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 || id >= len(movies) {
		respondWithError(w, http.StatusNotFound, "Movie not found")
		return
	}

	writeJSONResponse(w, http.StatusOK, movies[id])
}

// CreateMovie creates a new movie.
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	var newMovie Movie
	if err := json.NewDecoder(r.Body).Decode(&newMovie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request format")
		return
	}

	movies = append(movies, newMovie)

	writeJSONResponse(w, http.StatusCreated, newMovie)
}

// UpdateMovie updates an existing movie by ID.
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 || id >= len(movies) {
		respondWithError(w, http.StatusNotFound, "Movie not found")
		return
	}

	var updatedMovie Movie
	if err := json.NewDecoder(r.Body).Decode(&updatedMovie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request format")
		return
	}

	movies[id] = updatedMovie

	writeJSONResponse(w, http.StatusOK, updatedMovie)
}

// DeleteMovie deletes a movie by ID.
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 || id >= len(movies) {
		respondWithError(w, http.StatusNotFound, "Movie not found")
		return
	}

	movies = append(movies[:id], movies[id+1:]...)

	w.WriteHeader(http.StatusNoContent)
}
