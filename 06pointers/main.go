package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")

	var num1 = 25

	// copying the value
	duplicateNum1 := num1
	duplicateNum1 += 2
	fmt.Println("value of variable num1 is:", num1)

	// passing memory
	var pointerRef = &num1
	*pointerRef += 2
	fmt.Println("value of variable duplicateNum1 is:", duplicateNum1)
	fmt.Println("value of variable num1 is:", num1)

	fmt.Println("reference of variable pointerRef is:", pointerRef)
	fmt.Println("value of variable pointerRef is:", *pointerRef)

}
