package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	tom := User{"Tom", "Tom@go.dev", 12, true}
	fmt.Println("The user tom is: ", tom)
	fmt.Printf("The user tom is: %+v\n", tom)
	fmt.Printf("The name is %v, and the email is %+v\n", tom.Name, tom.Email)
	// Functions inside a class is specifically called as methods
	tom.GetStatus()
	tom.NewEmail()
	fmt.Printf("The name is %v, and the email is %+v\n", tom.Name, tom.Email)
}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

// these functions goes in User struct and called as methods
func (u User) GetStatus() {
	fmt.Println("The status of the user is: ", u.Status)
}

func (u User) NewEmail() {
	u.Email = "newtom@go.dev"
	fmt.Println("The new email is: ", u.Email)
}

// In Golang most of the times a reference of the variable is created and passed rather than using the actual variable in case of built-in data types
// hence, changing the value has no effect on the actual variable.
// That's where pointers play important role
// Execute and see why these comments are here
