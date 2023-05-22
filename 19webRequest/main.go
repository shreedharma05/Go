package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Web requests in Golang")

	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	fmt.Printf("The type of response is: %T\n", response)
	fmt.Println("The response is: ", response)
	fmt.Println("The response body is: ", response.Body)

	// Reads all data and returns them in byte codes
	dataBytes, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
	}

	// converts byte codes into string
	content := string(dataBytes)
	fmt.Println(content)
}
