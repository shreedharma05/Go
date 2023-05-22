package routes

import (
	"github.com/gorilla/mux"
	controllers "github.com/shreedharma05/go/mongodbapi/Controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controllers.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controllers.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controllers.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/api/movies", controllers.DeleteMovies).Methods("DELETE")

	return router
}
