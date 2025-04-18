package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	fmt.Println("value of dice is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("dice value is 1 and you can open")
	case 2:
		fmt.Println("dice value is 2 and you can move 2 spots")
	case 3:
		fmt.Println("dice value is 3 and you can move 3 spots")
		fallthrough
	case 4:
		fmt.Println("dice value is 4 and you can move 4 spots")
	case 5:
		fmt.Println("dice value is 5 and you can move 5 spots")
	case 6:
		fmt.Println("dice value is 6 and you can move 2 spots and roll dice again")
	default:
		fmt.Println("invalid number")
	}
}
