package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Random in Go")

	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(5))
}
