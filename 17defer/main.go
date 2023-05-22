package main

import "fmt"

// All Defer decorated statements are executed at the end of the scope (scope refered as a function)
// All Defered statements are executed in LIFO order (last in first out)
func main() {
	fmt.Println("Defer in Golang")
	defer fmt.Println("World")
	defer fmt.Println("Hello")
	defer fmt.Println("One")
	// for i := 0; i < 5; i++ {
	// 	defer fmt.Println(i)
	// }
	fmt.Println("Um, Hello")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
