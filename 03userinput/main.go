package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	str := "Hello, please enter your name:"
	fmt.Println(str)

	reader := bufio.NewReader(os.Stdin)

	// comma ok || comma err syntax

	input, _ := reader.ReadString('\n')
	fmt.Println("Hi, ", input)
	fmt.Printf("type of this input is %T", input)
}
