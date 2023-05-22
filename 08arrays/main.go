package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to arrays in Golang")

	var fruitsList [4]string
	fmt.Printf("The type is: %T\n", fruitsList)
	fruitsList[0] = "pineapple"
	fruitsList[1] = "banana"
	fruitsList[3] = "mango"
	fmt.Println("Fruits list is: ", fruitsList)
	fmt.Println("Fruits list is: ", len(fruitsList))

	var vegList = [5]string{"carrot", "beetroot", "brocoli"}
	fmt.Println("Fruits list is: ", vegList)
	fmt.Println("Fruits list is: ", len(vegList))

}
