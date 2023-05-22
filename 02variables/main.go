package main

import "fmt"

const loginToken = "myLoginToken"

const LoginToken = "myloginToken" // Public - accessible outside this package or simply exported (variables starting with uppercase)

func main() {

	// -----------------------------------------------------------------

	// Go creates a copy when assigning to a variable depends on the type of value being assigned,
	// with built-in types being copied and composite types being assigned a new reference to the original data structure.

	// built-in types such as int, float64, bool, string, and arrays, assigning one variable to another creates a copy of the underlying value.
	// x := 10
	// y := x // creates a copy of x's value, so y now contains 10 independent from x

	// -----------------------------------------------------------------

	var username string = "Shreedharma"
	fmt.Println(username)
	fmt.Printf("your username is a type of %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("your isLoggedIn is a type of %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("your smallVal is a type of %T \n", smallVal)

	var smallFloat float32 = 255.244434444444454444
	fmt.Println(smallFloat)
	fmt.Printf("your smallFloat is a type of %T \n", smallFloat)

	// default value and aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("your anotherVariable is a type of %T \n", anotherVariable)

	var yetAnotherVariable string
	fmt.Println(yetAnotherVariable)
	fmt.Printf("your yetAnotherVariable is a type of %T \n", yetAnotherVariable)

	// implicit type - lexer sets the type
	var numberOfUser = 324
	fmt.Println(numberOfUser)
	fmt.Printf("your numberOfUser is a type of %T \n", numberOfUser)

	// Shorthand for declaration and assignment (varless syntax) - but only works inside a function not outside
	jsonToken := "thisIsMyToken"
	fmt.Println(jsonToken)
	fmt.Printf("your jsonToken is a type of %T \n", jsonToken)

	fmt.Println(loginToken)
	fmt.Printf("your loginToken is a type of %T \n", loginToken)
}

// jwtToken := "thisIsNotMyToken" - Not gonna work.
