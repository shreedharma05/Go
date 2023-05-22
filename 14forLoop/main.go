package main

import "fmt"

func main() {
	fmt.Println("Loops in Golang")
	days := []string{"Sunday", "Monday", "Wednesday", "Thursday", "Friday", "Saturday"}

	// Classic for loop
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	fmt.Printf("\n")

	// Introducing range
	for j := range days {
		fmt.Println(days[j])
	}

	fmt.Printf("\n")

	// range to next level
	for idx, val := range days {
		fmt.Printf("The index is %v, and the value is %v\n", idx, val)
	}
	// for _, val := range days {
	// 	fmt.Printf("The index is %v, and the value is %v\n", idx, val)
	// }

	fmt.Printf("\n")

	value := 1
	// introducing break, continue, goto, and kind of while loop syntax
	for value <= 10 {
		if value == 2 {
			goto lco
		}
		if value == 5 {
			// break
			value++
			continue
		}
		fmt.Println(value)
		value++
	}
lco:
	fmt.Println("This is the goto")
}
