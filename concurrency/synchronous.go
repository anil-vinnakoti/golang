package main

import (
	"fmt"
	"time"
)

func Synchronous() {

	start := time.Now()

	ExecuteTask("Task 1", 2)
	ExecuteTask("Task 2", 2)
	ExecuteTask("Task 3", 2)

	fmt.Printf("Total time: %v\n", time.Since(start))
}
