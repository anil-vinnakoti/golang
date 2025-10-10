package main

import "fmt"

func CloseChannel() {
	ch := make(chan int, 2)

	ch <- 10
	ch <- 20
	close(ch)

	for i := 0; i < 3; i++ {
		val, ok := <-ch
		fmt.Println("Value: ", val, "open", ok)
	}
}

func ChannelRange() {
	ch := make(chan string, 3)

	ch <- "one"
	ch <- "two"
	ch <- "three"
	close(ch)

	for msg := range ch {
		fmt.Println("recieved: ", msg)
	}

	fmt.Println("channel closed, loop ended")
}

// Always close a channel from the sender’s side when all values are sent. Without closing, the range loop would block forever, waiting for more values that will never arrive.
// Never send after close
// Returns remaining values, then zero value
// Stops automatically after close
