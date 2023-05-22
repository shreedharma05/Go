package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in Golang")

	// ----------------------------------------------------------------------------------------

	// Go creates a copy when assigning to a variable depends on the type of value being assigned,
	// with built-in types being copied and composite types being assigned a new reference to the original data structure.

	// slice1 := []int{1, 2, 3}
	// slice2 := slice1 // creates a new reference to the underlying data
	// slice1[0] = 4 // modifies the underlying data through slice1
	// fmt.Println(slice2) // Output: [4 2 3]

	// composite types such as slices, maps, and channels, the behavior is more complex.
	// When assigning one variable to another, Go creates a new reference to the underlying data structure, but it does not create a copy of the data itself.
	// This means that changes made to the data through one variable will be visible through the other variable, since they both refer to the same data structure.

	// ----------------------------------------------------------------------------------------

	// if size is not explicitly given then it is a slice
	var fruitList = []string{"Grapes"}
	fmt.Printf("The type of fruitList is %T\n", fruitList)
	fruitList = append(fruitList, "Banana", "Mango", "Pineapple")
	fmt.Println("The fruitList is: ", fruitList)
	// fruitList = append(fruitList[1:3])
	fmt.Println("The sliced fruitList is: ", fruitList)

	// Declare using make function

	var highScores = make([]int, 4)
	highScores[0] = 123
	highScores[1] = 345
	highScores[2] = 567
	highScores[3] = 789
	// In case of make function you explicitly declare the size of the slice, so you can only add additional elements by append function
	fmt.Println("The highScores slice is: ", highScores)
	highScores = append(highScores, 1245, 898, 999)
	fmt.Println("The new highScores is: ", highScores)
	// sort
	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println("The sorted order of highScores is: ", highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// Remove an element from slice
	var courses = []string{"Reactjs", "Javascript", "Swift", "React-Native", "Ruby"}
	fmt.Println("The courses slice is: ", courses)
	var index int = 1
	// append function only individual elements to add, if you want to append a slice into a slice then use ellipsis(...).
	courses = append(courses[:index], courses[index+1:]...) // ... means all the elements inside this slice should be appended individually
	fmt.Println("The updated courses slice is: ", courses)
}
