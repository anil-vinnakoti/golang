package main

import "fmt"

func main() {
	num := 25

	// copying value
	copiedNum := num
	copiedNum += 5

	fmt.Println("copied num", copiedNum) // it is 30
	fmt.Println("actual num", num)       // ===> its 25

	// passing memory reference to the variable
	newNum := &num
	*newNum += 5

	fmt.Println("copied num", newNum) // it give the memory reference
	fmt.Println("actual num", num)    // ===> its 30
	fmt.Println("new num", *newNum)   // ===> its 30
}

// * Note *
// use & to pass the memory reference
// user * to access the value from the memory reference
