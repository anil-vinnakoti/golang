package main

import (
	"fmt"
	"net/url" // Tip: net/url provides tools to parse and construct URLs
)

// Define a sample URL with query parameters and a custom port
const myURL string = "https://jsonplaceholder.typicode.com:3000/comments?postId=1&userId=5"

func main() {
	fmt.Println("handling URLs in GO")
	fmt.Println(myURL)

	// ------------------------------------------
	// Parse the URL string into a *url.URL struct
	// Tip: url.Parse returns a parsed URL or error
	// ------------------------------------------
	res, err := url.Parse(myURL)
	if err != nil {
		panic(err)
	}

	// Access parts of the parsed URL
	fmt.Println(res.Scheme)   // Output: https
	fmt.Println(res.Host)     // Output: jsonplaceholder.typicode.com:3000
	fmt.Println(res.Path)     // Output: /comments
	fmt.Println(res.Port())   // Output: 3000
	fmt.Println(res.RawQuery) // Output: postId=1&userId=5

	// ------------------------------------------
	// Extract query parameters as a map
	// Tip: res.Query() returns url.Values (map[string][]string)
	// ------------------------------------------
	qParams := res.Query()
	fmt.Printf("the type of query params are: %T\n", qParams)

	// Access values by key (returns a slice of strings)
	fmt.Println(qParams["postId"]) // Output: [1]

	// Iterate through all query parameters
	for key, val := range qParams {
		fmt.Printf("param key is %v and value is %v:\n", key, val)
	}

	// ------------------------------------------
	// Construct a new URL using url.URL struct
	// Tip: &url.URL{} creates a pointer to a URL struct,
	//      required to call pointer receiver methods like .String()
	// Note: RawPath is used for escaped path segments, not query
	// Use RawQuery for query parameters
	// ------------------------------------------
	partsOfURL := &url.URL{
		Scheme:   "https",
		Host:     "localhost",
		Path:     "users",
		RawQuery: "user=1&post=3", // Tip: Use RawQuery for query parameters
	}

	// Convert URL object to string
	anotherURL := partsOfURL.String()
	fmt.Println(anotherURL) // Output: https://localhost/users?user=1&post=3
}
