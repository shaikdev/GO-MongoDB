package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
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
	defer r.Body.Close()
	getMovies, count := service.GetAllMovies()
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"message":    responses.GET_ALL_MOVIES_SUCCESSFULLY,
		"status":     "Success",
		"statusCode": http.StatusOK,
		"data":       getMovies,
		"count":      count,
	})
}

func GetMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	defer r.Body.Close()
	params := mux.Vars(r)
	movieId := params["id"]
	response, err := service.GetMovieById(movieId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"message":    responses.MOVIE_GET_FAILED,
			"status":     "Failed",
			"statusCode": http.StatusNotFound,
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"message":    responses.MOVIE_GET_SUCCESS_SUCCESSFULLY,
		"status":     "Success",
		"statusCode": http.StatusOK,
		"data":       response,
	})
}

func UpdateMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	defer r.Body.Close()
	params := mux.Vars(r)
	movieId := params["id"]
	// decode
	var movie models.Movie
	json.NewDecoder(r.Body).Decode(&movie)
	response := service.UpdateMovie(movieId, movie)
	if response == 0 {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Failed to edit movie",
		})
		return
	}
	_ = json.NewEncoder(w).Encode("Updated successfully")

}

func DeleteMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	movieId := params["id"]
	response := service.DeleteMovieById(movieId)
	if response == 0 {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"message":    responses.MOVIE_DELETED_FAILED,
			"status":     "Failed",
			"statusCode": http.StatusBadRequest,
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"message":    responses.MOVIE_DELETED_SUCCESSFULLY,
		"status":     "Success",
		"statusCode": http.StatusOK,
	})
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	response := service.DeleteAllMovies()
	if response == 0 {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"message":    responses.MOVIE_DELETE_FAILED,
			"statusCode": http.StatusBadRequest,
			"status":     "Failed",
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"message":    responses.MOVIE_DELETED_SUCCESSFULLY,
		"statusCode": http.StatusOK,
		"status":     "Success",
	})
}
