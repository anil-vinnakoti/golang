package main

import (
	"fmt"
	"time"
)

func Asynchronous() {

	start := time.Now()

	go ExecuteTask("Task 1", 2)
	go ExecuteTask("Task 2", 2)
	go ExecuteTask("Task 3", 2)

	time.Sleep(3 * time.Second)

	fmt.Printf("Total time: %v\n", time.Since(start))
}
