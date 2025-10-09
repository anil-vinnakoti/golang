package main

import (
	"fmt"
	"sync"
	"time"
)

func WaitGroups() {
	start := time.Now()

	wg := sync.WaitGroup{}

	wg.Add(3)
	go func() {
		defer wg.Done()
		ExecuteTask("Task 1", 2)
	}()

	go func() {
		defer wg.Done()
		ExecuteTask("Task 2", 2)
	}()
	go func() {
		defer wg.Done()
		ExecuteTask("Task 3", 2)
	}()

	wg.Wait()

	fmt.Printf("Total time: %v\n", time.Since(start))
}
