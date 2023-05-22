package main

import "fmt"

func main() {
	fmt.Println("Functions in Golang")
	// You cannot have functions inside a function
	// Immedietly invoked functions are available
	// anonymous functions are also avilable
	greeter()
	fmt.Println("The sum is: ", adder(2, 2))
	sum, message := proAdder(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Printf("The total is %v, and the message recieved is %v", sum, message)
}

func greeter() {
	fmt.Println("Hello from Golang")
}
func adder(a int, b int) int {
	return a + b
}

// takes any number of parameters
func proAdder(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, "proAdder executed"
}
