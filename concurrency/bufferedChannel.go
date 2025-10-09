package main

import (
	"fmt"
)

func BufferedChannel() {
	ch := make(chan string, 2)

	go func() {
		ch <- "msg 1"
		fmt.Println("Sent msg 1")

		ch <- "msg 2"
		fmt.Println("Sent msg 2")

		ch <- "msg 3" // 🧱 will block until a receive happens
		fmt.Println("Sent msg 3")
	}()

	fmt.Println("Receiving...")
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
