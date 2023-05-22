package main

import (
	"fmt"
	"net/url"
)

const myURL = "https://lco.dev:3000/learn?coursename=nestjs&paymentid=123lpd456"

func main() {
	fmt.Println("URLs in Golang")
	fmt.Printf("The URL is: %v \nand the type of URL is: %T\n", myURL, myURL) // Returns a myURL and type is string

	result, _ := url.Parse(myURL)

	fmt.Printf("The URL is: %v \nand the type of URL is: %T\n", result, result) // Returns a struct containing components of the url where data can be extracted,
	// Such as:
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	qParams := result.Query() // Returns a key value pair (Maps) named url.values of the params present in the url
	fmt.Printf("The params are: %v, \n and the type of params is: %T \n", qParams, qParams)
	for _, value := range qParams {
		fmt.Printf("The parameter is: %v\n", value)
	}

	// Creating a url with URL struct in url package
	partsOfURL := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=tom",
	}
	fmt.Println("1: ", partsOfURL)
	fmt.Println("2: ", *partsOfURL)
	newUrl := partsOfURL.String() // For url you are supposed to use only this syntax to convert to string and not string(partsOfURL)

	fmt.Println("The new URL is: ", newUrl)
}
