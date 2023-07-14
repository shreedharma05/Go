package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	//
	tom := User{"Tom", "Tom@go.dev", 12, true}
	fmt.Println("The user tom is: ", tom)
	fmt.Printf("The user tom is: %+v\n", tom)
	fmt.Printf("The name is %v, and the email is %+v\n", tom.Name, tom.Email)

	// personPtr is a pointer of Person struct, passed by reference
	personPtr := &Person{
		Name:       "Bob",
		Age:        25,
		Occupation: "Teacher",
	}

	//
	fmt.Println(*personPtr)           // Output:
	fmt.Println(personPtr.Age)        // Output: 25
	fmt.Println(personPtr.Occupation) // Output: Teacher

}

// is exported if starts with uppercase
type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

type Person struct {
	Name       string
	Age        int
	Occupation string
}
