// exported packages can be accessed via package name
// any variables or functions starts with uppercase are exported
package main // exported as 01hello/main

// semicolons are mandatory but lexer takes care of it

import "fmt"

func main() {
	// Print()
	// Println()
	// Printf()
	fmt.Println("Um, Hello!")
	fmt.Println("Um, Hello!")
}
