package main

import (
	"fmt"
	"log"
	"net/http"

	routes "github.com/shreedharma05/go/mongodbapi/Routes"
)

func main() {
	fmt.Println("MongoDB API")
	r := routes.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":8000", r)) // Listen at port 8000 and serve router "r" at that port
	fmt.Println("Listening at port 8000...")
}
