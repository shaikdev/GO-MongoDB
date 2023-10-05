package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shaikdev/GO-MongoDB/helper"
	"github.com/shaikdev/GO-MongoDB/models"
	"github.com/shaikdev/GO-MongoDB/responses"
	"github.com/shaikdev/GO-MongoDB/service"
)

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	helper.Header(w, "POST")
	defer r.Body.Close()
	var movie models.Movie
	json.NewDecoder(r.Body).Decode(&movie)
	if movie.BodyCheckForMovie() {
		helper.ResponseErrorSender(w, responses.MOVIE_FIELD_IS_REQUIRED, "Failed", http.StatusBadRequest)
		return
	}
	response, createErr := service.CreateMovie(movie)
	if createErr != nil {
		helper.ResponseErrorSender(w, responses.MOVIE_CREATED_FAILED, "Failed", http.StatusBadRequest)
		return
	}
	helper.ResponseSuccessSender(w, responses.MOVIE_CREATED_SUCCESSFULLY, "Success", http.StatusCreated, response)
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	helper.Header(w, "GET")
	defer r.Body.Close()
	getMovies, count := service.GetAllMovies()
	helper.ResponseSuccessSenderWithCount(w, responses.GET_ALL_MOVIES_SUCCESSFULLY, "Success", http.StatusOK, getMovies, count)
}

func GetMovieById(w http.ResponseWriter, r *http.Request) {
	helper.Header(w, "GET")
	defer r.Body.Close()
	params := mux.Vars(r)
	movieId := params["id"]
	response, err := service.GetMovieById(movieId)
	if err != nil {
		helper.ResponseErrorSender(w, responses.MOVIE_GET_FAILED, "Failed", http.StatusNotFound)
		return
	}
	helper.ResponseSuccessSender(w, responses.MOVIE_GET_SUCCESS_SUCCESSFULLY, "Success", http.StatusOK, response)
}

func UpdateMovieById(w http.ResponseWriter, r *http.Request) {
	helper.Header(w, "PUT")
	defer r.Body.Close()
	params := mux.Vars(r)
	movieId := params["id"]
	// decode
	var movie models.Movie
	json.NewDecoder(r.Body).Decode(&movie)

	err := service.UpdateMovie(movieId, movie)
	if err != nil {
		helper.ResponseErrorSender(w, responses.MOVIE_EDIT_FAILED, "Failed", http.StatusBadRequest)
		return
	}

	updatedMovie, getMovieErr := service.GetMovieById(movieId)
	if getMovieErr != nil {
		helper.ResponseErrorSender(w, responses.MOVIE_GET_FAILED, "Failed", http.StatusNotFound)
		return
	}
	helper.ResponseSuccessSender(w, responses.MOVIE_EDIT_SUCCESSFULLY, "Success", http.StatusOK, updatedMovie)

}

func DeleteMovieById(w http.ResponseWriter, r *http.Request) {
	helper.Header(w, "DELETE")

	params := mux.Vars(r)
	movieId := params["id"]
	response := service.DeleteMovieById(movieId)
	if response == 0 {
		helper.ResponseErrorSender(w, responses.MOVIE_DELETE_FAILED, "Failed", http.StatusBadRequest)
		return
	}
	helper.ResponseSuccessSenderWithoutData(w, responses.MOVIE_DELETED_SUCCESSFULLY, "Success", http.StatusOK)
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	helper.Header(w, "DELETE")
	response := service.DeleteAllMovies()
	if response == 0 {
		helper.ResponseErrorSender(w, responses.MOVIE_DELETE_FAILED, "Failed", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	helper.ResponseSuccessSenderWithoutData(w, responses.MOVIE_DELETED_SUCCESSFULLY, "Success", http.StatusOK)
}
