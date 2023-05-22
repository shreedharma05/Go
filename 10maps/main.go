package main

import "fmt"

func main() {

	fmt.Println("Maps in golang")
	// key value pair
	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["GO"] = "Golang"
	languages["JV"] = "JAVA"
	fmt.Println("The languages map is: ", languages)
	fmt.Printf("The type of languages is: %T\n", languages)

	// delete from map
	delete(languages, "PY")
	fmt.Println("The updated map is: ", languages)

	// Loops in maps
	for key, value := range languages {
		fmt.Printf("For the key %v, The value is %v\n", key, value)
	}
}
