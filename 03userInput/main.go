package main

// importing multiple packages
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// creates a NewReader instance which is capable of reading input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Burger")

	// Reads until \n (nextline)
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating ", input)
	// %T gives type of the instance
	fmt.Printf("the type of rating is %T ", input)

	// fmt.Println()
}
