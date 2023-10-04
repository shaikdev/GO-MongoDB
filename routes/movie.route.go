package routes

import (
	"github.com/gorilla/mux"
	"github.com/shaikdev/GO-MongoDB/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/movie", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/api/v1/movies", controllers.GetAllMovies)
	return router

}
