package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		json.NewEncoder(w).Encode("Path not available")
	}
	json.NewEncoder(w).Encode("Um, Hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("Name")
	age := r.FormValue("Age")
	address := r.FormValue("Address")

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %s\n", age)
	fmt.Printf("Address %s\n", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("The server is starting at port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
