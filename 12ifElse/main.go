package main

import "fmt"

func main() {
	var number = 45
	result := ""
	if number < 10 {
		result = "Number less than 10"
	} else if number > 10 {
		result = "Number greater than 10"
	} else {
		result = "Number exactly equal to 10"
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("The Number is even")
	} else {
		fmt.Println("The Number is odd")
	}

	if x := 6; x < 5 {
		fmt.Println("The value of x is less than 5")
	} else if x == 5 {
		fmt.Println("The value of x is 5")
	} else {
		fmt.Println("The value of x is greater than 5")
	}
}
