package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Switch case in Golang")

	// seeding a number using time so we always get a unique number
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(6) + 1

	fmt.Println("The number is: ", number)

	switch number {
	case 1:
		fmt.Println("The number is 1 and you can open")
	case 2:
		fmt.Println("You can move two spot")
	case 3:
		fmt.Println("You can move three spot")
		fallthrough
	case 4:
		fmt.Println("You can move four spot")
		fallthrough
	case 5:
		fmt.Println("You can move five spot")
	case 6:
		fmt.Println("You can move six spot and roll one more time")
	}
}
