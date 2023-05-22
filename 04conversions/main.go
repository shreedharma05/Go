package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	// by default this rating variable include \n in the end (includes an empty line)
	rating, _ := reader.ReadString('\n')

	fmt.Println("Thank you for rating ", rating)

	// TrimSpace removes empty spaces (function in strings package)
	// strconv is package which has ParseFloat function used to extract the string as float of 64 bit size
	newRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to the rating ", newRating+1)
	}
}
