package main

import "fmt"

func main() {
	// Simulate user login count
	loginCount := 25

	// Declare a variable to store result based on loginCount
	var result string

	// Tip: Use if-else chains to handle multiple conditional branches
	if loginCount < 15 {
		result = "normal user"
	} else if loginCount == 25 {
		result = "logged in 25 times"
	} else {
		result = "regular user"
	}

	// Output the result
	fmt.Println(result)

	// ---------------------------------------
	// Using short variable declaration inside if
	// ---------------------------------------
	// Tip: The variable 'num' is scoped only within this if-else block
	if num := 5; num < 10 {
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("num is greater than 10")
	}

	// ---------------------------------------
	// Go idiom: Error handling placeholder
	// ---------------------------------------
	// Tip: Go doesn't use try-catch; instead, functions often return (value, error)
	// This block is commonly used when calling such functions:
	//
	// if err != nil {
	//     fmt.Println("An error occurred:", err)
	//     return
	// }
	//
	// Example:
	// data, err := ioutil.ReadFile("file.txt")
	// if err != nil {
	//     log.Fatal(err)
	// }

}
