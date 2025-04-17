package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome to arrays")

	var list [3]string

	list[0] = "a"
	list[2] = "c"

	// you will see a long space between index values of 0 and 2, because 2 is empty and its a string
	fmt.Println("list", list)

	fmt.Println("list length", len(list)) //3

	var list2 = [4]int{1, 2, 3, 4}
	fmt.Println("integers list", list2)
	fmt.Println("list2 length", len(list2)) //4
}
