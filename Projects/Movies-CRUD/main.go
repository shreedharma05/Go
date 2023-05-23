package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Distributor string    `json:"distributor"`
	Director    *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var Movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, val := range Movies {
		if val.Id == params["id"] {
			Movies = append(Movies[:idx], Movies[idx+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, val := range Movies {
		if val.Id == params["id"] {
			json.NewEncoder(w).Encode(Movies[idx])
			break
		}
	}
	json.NewEncoder(w).Encode(Movies)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)
	rand.Seed(time.Now().UnixNano())
	movie.Id = strconv.Itoa(rand.Intn(1000))
	Movies = append(Movies, movie)
	json.NewEncoder(w).Encode(Movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, val := range Movies {
		if val.Id == params["id"] {
			Movies = append(Movies[:idx], Movies[idx+1:]...)
			var movie Movie
			json.NewDecoder(r.Body).Decode(&movie)
			movie.Id = params["id"]
			Movies = append(Movies, movie)
		}
	}
	json.NewEncoder(w).Encode(Movies)
}

func main() {
	r := mux.NewRouter()

	Movies = append(Movies, Movie{Id: "1", Title: "Oppenheimer", Distributor: "Universal pictures", Director: &Director{Firstname: "Christopher", Lastname: "Nolan"}})
	Movies = append(Movies, Movie{Id: "2", Title: "Barbie", Distributor: "Warner Bros", Director: &Director{Firstname: "Greta", Lastname: "Gerwig"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movie", createMovie).Methods("POST")
	r.HandleFunc("/movie/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Server is running at port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))

}
