package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Welcome to Time tutorial of Golang")

	// inbuilt package for time
	presentTime := time.Now().Format("Monday 15:03:05")
	fmt.Println(presentTime)

	// Format function takes a format in which time should be returned (values inside format are fixed, these exact values are only allowed)
	// fmt.Println(presentTime.Format("01-02-2006 Monday 15:03:05"))

	// Date function is to create a new date (usally for setting deadlines)
	createdDate := time.Date(2024, time.January, 05, 00, 00, 00, 0, time.UTC)
	fmt.Println(createdDate)

	fmt.Println(createdDate.Format("Monday 15:03:05"))
}
