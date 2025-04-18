package main

import "fmt"

func main() {
	fmt.Println("welcome to functions")

	// ----------------------------------------
	// Calling a simple function
	// ----------------------------------------
	wish()

	// ----------------------------------------
	// Calling a function with two int parameters
	// ----------------------------------------
	result := adder(5, 8)
	fmt.Println("adder result is:", result)

	// ----------------------------------------
	// Calling a variadic function
	// ----------------------------------------
	// Tip: Use variadic functions when the number of arguments can vary
	result = addMultipleValues(1, 2, 8, 9)
	fmt.Println("multiple values result is:", result)

	// ----------------------------------------
	// Receiving multiple return values
	// ----------------------------------------
	// Tip: Go supports returning multiple values from a function
	num, str := returnMultipleValues()
	fmt.Println("multiple return values:", num, str)

	// ----------------------------------------
	// Anonymous function example
	// ----------------------------------------
	// Tip: You can define and call anonymous functions inside other functions
	// Tip: You CANNOT define a named function inside another function in Go
	greet := func(name string) {
		fmt.Printf("Hello, %s!\n", name)
	}
	greet("Anil")

	// ❌ Invalid: This would cause a compile-time error
	// func innerFunction() {
	//     fmt.Println("This is not allowed!")
	// }
}

// ----------------------------------------
// Function to add two integers
// Tip: Go requires specifying types for all parameters
// ----------------------------------------
func adder(a int, b int) int {
	return a + b
}

// ----------------------------------------
// Variadic function to sum any number of integers
// Tip: '...int' lets you pass zero or more int arguments
// ----------------------------------------
func addMultipleValues(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

// ----------------------------------------
// Function returning multiple values
// Tip: Use this for returning result + error, or multiple data pieces
// ----------------------------------------
func returnMultipleValues() (int, string) {
	return 1, "hello"
}

// Simple function to print a message
func wish() {
	fmt.Println("hello")
}
