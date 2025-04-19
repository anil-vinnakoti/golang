package main

import (
	"fmt"
)

func main() {
	// Tip: defer statements are executed in **reverse order (LIFO)** when the function exits
	defer fmt.Println("world")
	defer fmt.Println("one")
	defer fmt.Println("two")

	fmt.Println("hello") // This prints first

	myDefer()
}

func myDefer() {
	// Tip: Each deferred call captures the value of 'i' at the time it is deferred
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
	// Output will be:
	// 3
	// 2
	// 1
	// 0
}
