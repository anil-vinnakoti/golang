package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("welcome to web requests")
	const getUrl = "http://localhost:5000/get"
	const postUrl = "http://localhost:5000/post"
	const postFormUrl = "http://localhost:5000/postform"
	// performGetRequest(getUrl)
	performPostFormRequest(postFormUrl)

}

// performGetRequest sends a GET request to the provided endpoint
func performGetRequest(endpoint string) {
	// Tip: http.Get returns a response and an error
	res, err := http.Get(endpoint)
	handleError(err)
	defer res.Body.Close() // Tip: Always close response body to avoid memory leaks

	fmt.Println("Status code:", res.StatusCode)
	fmt.Println("Content length is:", res.ContentLength)

	// -------------------------------------------
	// Read response body using io.ReadAll
	// Tip: io.ReadAll reads all bytes from the response body
	// -------------------------------------------
	content, err := io.ReadAll(res.Body)
	handleError(err)

	// ------------------------------
	// Option 1: Directly convert bytes to string
	// Tip: Simple and quick for small data
	// ------------------------------
	byteCount := len(content)
	fmt.Println("Byte count is:", byteCount)
	fmt.Println("Content using string():", string(content))

	// ------------------------------
	// Option 2: Use strings.Builder
	// Tip: More efficient when working with large or dynamic strings
	// ------------------------------
	var resString strings.Builder
	builderByteCount, _ := resString.Write(content)
	fmt.Println("Byte count (using Builder):", builderByteCount)
	fmt.Println("Content using strings.Builder:", resString.String())
}

func performPostJsonRequest(endpoint string) {

	//fake json payload
	requestBody := strings.NewReader(`{
		"name": "Anil",
		"age": "27"
	}`)

	res, err := http.Post(endpoint, "application/json", requestBody)

	handleError(err)

	defer res.Body.Close()

	content, err := io.ReadAll(res.Body)
	handleError(err)

	fmt.Println(string(content))

}

func performPostFormRequest(endpoint string) {
	//fake form data
	data := url.Values{}
	data.Add("firstName", "Anil")
	data.Add("lastName", "Vinnakoti")
	data.Add("email", "anil.vinnakoti@gmail.com")

	res, err := http.PostForm(endpoint, data)
	handleError(err)
	defer res.Body.Close()

	content, err := io.ReadAll(res.Body)
	handleError(err)
	fmt.Println(string(content))

}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
