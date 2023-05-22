package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hey, All!")
	var slice1 = []int{5, 4}
	var slice2 = []int{1, 3, 2, 4, 5}
	fmt.Println(prob_1(slice1, slice2))
}

func prob_1(slice1 []int, slice2 []int) int {
	var N int = slice1[0]
	X := slice1[1]
	sort.Ints(slice2)
	L := len(slice2)
	if X <= N {
		if slice2[L-X] == slice2[L-X-1] {
			return -1
		} else {
			return slice2[L-X]
		}
	} else {
		return -1
	}
}
