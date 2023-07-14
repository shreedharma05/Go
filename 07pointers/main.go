package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to Pointers in Golang")

	// Not all passing of pointers in Go is equivalent to passing by reference.
	// For example, when you pass a pointer to a variable to a function, you are passing a copy of the memory address of the variable,
	// not a reference to the variable itself. If the function modifies the pointer to point to a different value, that change will not be visible outside the function.

	// In general, it's more accurate to say that Go uses a combination of pass-by-value and pass-by-reference semantics.
	// When you pass a pointer to a function, you are passing a copy of the memory address by value,
	// but the function can use that copy to indirectly access and modify the original value in memory.

	// var ptr *int
	// fmt.Println("Value of ptr is", ptr)

	var x = 20
	var ptr = &x
	fmt.Println("value of actual pointer:", ptr)  // & - gives the memory address of the value
	fmt.Println("value of actual pointer:", *ptr) // * - gives the actual value
 fmt.Prinln(ptr +1)
	*ptr += 30
	fmt.Println("value of x:", x)

	if *ptr == x {
		fmt.Println(true)
	}

	// why pointer? - To perform actions exactly on the values and not on their copies
}
