package main

import "fmt"

// When pointers?
// 1. When you are changing the state of a variable
// 2. When passing by value is expensive - if the value carries heavy files, copying becomes expensive. Whereas a pointer is only 8 bytes

type User struct {
	Email string
	Name  string
	Age   int
}

// Pass by value - copy of User has been passed
func (u User) GetEmail_1() string {
	return u.Email
}

// Pass by refernece - reference to User has been passed
func (u *User) GetEmial_2() string {
	return u.Email
}

// Pass by value - updating value
func (u User) UpdateEmail_01(email string) {
	u.Email = email
}

// Pass by reference - updating value
func (u *User) UpdateEmail_02(email string) {
	u.Email = email
}

func main() {
	Joel := User{
		Email: "joel@go.dev",
		Name:  "Joel",
		Age:   22,
	}
	Joel.UpdateEmail_01("jo@go.dev")
	fmt.Println("While using pass by value: ", Joel.GetEmial_2())

	Joel.UpdateEmail_02("joe@go.dev")
	fmt.Println("While using pass by reference: ", Joel.GetEmail_1())
}
