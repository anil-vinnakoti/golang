package main

import "fmt"

// Define a struct type named User
type User struct {
	Name       string
	Email      string
	IsVerified bool
	Age        int
	Status     string
}

// Method with a value receiver (doesn't modify original struct)
func (u User) getStatus() {
	fmt.Println("user is", u.Status)
}

func main() {
	// Create an instance of User using a struct literal
	anil := User{
		Name:       "Anil",
		Email:      "anil.vinnakoti@gmail.com",
		IsVerified: true,
		Age:        27,
		Status:     "Online",
	}

	// Print the struct directly (default format)
	fmt.Println(anil)

	// Print with field names using %+v for better readability
	fmt.Printf("Anil's details are: %+v\n", anil)

	// Access and print individual fields
	fmt.Printf("Name is %v and email is %v\n", anil.Name, anil.Email)

	anil.getStatus()

	// ------------------------------------------
	// Attempting to update the email
	// Tip: This won't persist because value receivers receive a copy
	// ------------------------------------------
	anil.NewMail()
	fmt.Printf("After value receiver: Name is %v and email is %v\n", anil.Name, anil.Email)

	// ------------------------------------------
	// Now updating using pointer receiver
	// Tip: Use a pointer receiver to modify the original object
	// ------------------------------------------
	anil.UpdateMailWithPointer()
	fmt.Printf("After pointer receiver: Name is %v and email is %v\n", anil.Name, anil.Email)
}

// Method with a value receiver: changes won't reflect in the original object
func (u User) NewMail() {
	u.Email = "anil.ogl007@gmail.com"
	fmt.Println("[value receiver] email is changed to:", u.Email)
}

// Method with a pointer receiver: changes affect the original object
func (u *User) UpdateMailWithPointer() {
	u.Email = "anil.updated@gmail.com"
	fmt.Println("[pointer receiver] email is updated to:", u.Email)
}
