package main

import (
	"fmt"
	"time"
)

func ExecuteTask(name string, seconds int) {
	fmt.Printf("%v started\n", name)
	time.Sleep(time.Duration(seconds) * time.Second)
	fmt.Printf("%v finished\n", name)

}
