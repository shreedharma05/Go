package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const myGETURL = "http://localhost:8000/get"

// const myPOSTURL = "http://localhost:8000/post"
// const myPOSTFORMURL = "http://localhost:8000/postform"

func PerformGetRequest(theURL string) {
	response, err := http.Get(theURL)
	if err != nil {
		panic(err)
	}
	fmt.Println("The status of the response is: ", response.StatusCode)
	fmt.Println("The length of the content in response is: ", response.ContentLength)
	fmt.Println("The response is: ", response)

	fmt.Println("The response body is: ", response.Body)

	content, _ := io.ReadAll(response.Body)

	fmt.Println("The content is: ", content)

	// Beginners way
	// fmt.Println("The content as string is: ")
	// fmt.Println("The content is: ", string(content))

	// Better way
	var stringBuilder strings.Builder
	fmt.Println("The stringBuilder is: ", stringBuilder)
	byteCount, _ := stringBuilder.Write(content) // writes content to stringBuilder and returns the length of content (Purpose is to write content to stringBuilder)
	fmt.Println("The stringBuilder is: ", stringBuilder)
	fmt.Println("The byte count is: ", byteCount)
	fmt.Println("The content is: ", stringBuilder.String())
}

func PerformPostJSONRequest(theURL string) {
	// create fake payload
	requestBody := strings.NewReader(`{
		"coursename": "NestJS",
		"price": 10,
		"platform": "LCO"
	}`)

	response, err := http.Post(theURL, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	content, _ := io.ReadAll(response.Body)
	defer response.Body.Close()
	fmt.Println("The Content is: ", string(content))
}

func PerformPostFormRequest(theURL string) {
	data := url.Values{}
	data.Add("name", "Tom")
	data.Add("age", "10")
	data.Add("email", "tom@go.dev")

	response, err := http.PostForm(theURL, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println("The post form data is: ", string(content))
}

func main() {
	fmt.Println("Get request in Golang")
	PerformGetRequest(myGETURL)
	// PerformPostJSONRequest(myPOSTURL)
	// PerformPostFormRequest(myPOSTFORMURL)
}
