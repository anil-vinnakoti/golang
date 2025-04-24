package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var waitGroup sync.WaitGroup //pointer

var mut sync.Mutex //pointer

func main() {
	// go greeter("Hello")
	// greeter("World")

	websiteList := []string{
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		waitGroup.Add(1)
	}

	waitGroup.Wait()
	fmt.Println(signals)
}

// func greeter(str string) {

// 	for i := 0; i < 5; i++ {
// 		time.Sleep(1 * time.Second)
// 		fmt.Println(str)
// 	}
// }

func getStatusCode(endpoint string) {
	res, err := http.Get(endpoint)

	if err != nil {
		panic(err)
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()

		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
	defer waitGroup.Done()
}
