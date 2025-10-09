package main

import (
	"fmt"
	"time"
)

func UnbefferedChannel() {
	ch := make(chan string)

	go func() {
		fmt.Println("👷 Worker: preparing task...")
		ch <- "✅ Task done"
		fmt.Println("👷 Worker: sent the message")
	}()

	fmt.Println("🧍 Main: waiting for worker...")
	msg := <-ch
	fmt.Println("🧍 Main: received:", msg)

	time.Sleep(time.Second)
}
