package routes

import (
	"github.com/gorilla/mux"
	"github.com/shaikdev/GO-MongoDB/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/movie", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/api/v1/movies", controllers.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/v1/movie/{id}", controllers.GetMovieById).Methods("GET")
	router.HandleFunc("/api/v1/movie/{id}", controllers.UpdateMovieById).Methods("PUT")
	router.HandleFunc("/api/v1/movie/{id}", controllers.DeleteMovieById).Methods("DELETE")
	router.HandleFunc("/api/v1/movies", controllers.DeleteAllMovies).Methods("DELETE")
	return router

}
