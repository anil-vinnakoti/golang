package main

import (
	"fmt"
	"io"       // Tip: In Go 1.16+, use io.ReadAll instead of deprecated ioutil.ReadAll
	"net/http" // Tip: http package is used for making HTTP requests
)

const url = "https://jsonplaceholder.typicode.com/todos/1"

func main() {
	fmt.Println("web request")

	// Tip: http.Get performs a GET request and returns a response and error
	res, err := http.Get(url)
	if err != nil {
		panic(err) // Tip: panic stops the program immediately with an error message
	}

	// Tip: Always defer closing the response body to avoid memory/resource leaks
	defer res.Body.Close()

	fmt.Printf("response is type of %T\n", res) // Output the response type

	// Tip: io.ReadAll reads the entire response body into memory
	// Be cautious with large bodies — consider streaming or buffering if needed
	dataByte, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	// Convert the byte slice to a string to print readable content
	content := string(dataByte)
	fmt.Println(content)
}
