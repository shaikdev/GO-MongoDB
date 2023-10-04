package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/shaikdev/GO-MongoDB/models"
	"github.com/shaikdev/GO-MongoDB/responses"
	"github.com/shaikdev/GO-MongoDB/service"
)

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	defer r.Body.Close()
	var movie models.Movie
	json.NewDecoder(r.Body).Decode(&movie)
	if movie.BodyCheckForMovie() {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"message":    responses.MOVIE_FIELD_IS_REQUIRED,
			"status":     "Failed",
			"statusCode": http.StatusBadRequest,
		})
		return
	}
	response, createErr := service.CreateMovie(movie)
	if createErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"message":    responses.MOVIE_CREATED_FAILED,
			"status":     "Failed",
			"statusCode": http.StatusBadRequest,
		})
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"message":    responses.MOVIE_CREATED_SUCCESSFULLY,
		"status":     "Success",
		"statusCode": http.StatusCreated,
		"data":       response,
	})
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.WriteHeader(http.StatusOK)
	getMovies, count := service.GetAllMovies()
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"message":    responses.GET_ALL_MOVIES,
		"status":     "Success",
		"statusCode": http.StatusOK,
		"data":       getMovies,
		"count":      count,
	})
}
