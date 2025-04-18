package main

import "fmt"

// Define a struct type named User
type User struct {
	Name       string
	Email      string
	IsVerified bool
	Age        int
}

func main() {
	// Create an instance of User using a struct literal
	anil := User{
		Name:       "Anil",
		Email:      "anil.vinnakoti@gmail.com",
		IsVerified: true,
		Age:        27,
	}

	// Print the struct directly (default format)
	fmt.Println(anil)

	// Print with field names using %+v for better readability
	fmt.Printf("Anil's details are: %+v\n", anil)

	// Access and print individual fields
	fmt.Printf("Name is %v and email is %v\n", anil.Name, anil.Email)
}
