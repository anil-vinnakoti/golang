package main

import "fmt"

// constants
const pi = 3
const pi2 = 3.1415926535897932384626433832795
const firstName = "Anil Kumar"

func main() {
	var name string = "Anil"
	fmt.Println(name)
	fmt.Printf("Variable is of type: %T \n", name)

	var isMale bool = true
	fmt.Println(isMale)
	fmt.Printf("Variable is of type: %T \n", isMale)

	var age int = 255
	fmt.Println(age)
	fmt.Printf("Variable is of type: %T \n", age)

	var floatValue float32 = 255.2548766213546697533
	fmt.Println(floatValue)
	fmt.Printf("Variable is of type: %T \n", floatValue)

	//default values and some aliases
	var anotherIntVariable int
	fmt.Println(anotherIntVariable)
	fmt.Printf("Variable default value is: %T \n", anotherIntVariable)

	var anotherBoolVariable bool
	fmt.Println(anotherBoolVariable)
	fmt.Printf("Variable default value is: %T \n", anotherBoolVariable)

	var anotherStringVariable string
	fmt.Println(anotherStringVariable)
	fmt.Printf("Variable default value is: %T \n", anotherStringVariable)

	var anotherFloatVariable float64
	fmt.Println(anotherFloatVariable)
	fmt.Printf("Variable default value is: %T \n", anotherFloatVariable)

	//implicit type
	var implicitIntVariable = 100
	implicitIntVariable = 200 // This is valid because implicitIntVariable is of type int
	// implicitIntVariable = 300.5 // This will cause a compile-time error because implicitIntVariable is of type int
	// implicitIntVariable = "hello" // This will cause a compile-time error because implicitIntVariable is of type int
	fmt.Println(implicitIntVariable)
	fmt.Printf("Variable is of type: %T \n", implicitIntVariable)

	//short variable declaration
	implicitStringVariable := "Anil"
	implicitStringVariable = "Anil Kumar" // This is valid because implicitStringVariable is of type string
	// implicitStringVariable = 100 // This will cause a compile-time error because implicitStringVariable is of type string
	fmt.Println(implicitStringVariable)
	// can use walrus inside a method only and wont work outside a method

	fmt.Println("Pi value is: ", pi)
	fmt.Printf("Variable default value is: %T \n", pi)

}
