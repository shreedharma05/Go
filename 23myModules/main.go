package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to modules in Golang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("Get")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func greeter() {
	fmt.Println("Hey there mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Golang series</h1>"))
}
