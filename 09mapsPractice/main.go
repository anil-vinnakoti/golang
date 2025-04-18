package main

import "fmt"

func main() {
	// Create a map using make() — keys and values are strings
	langs := make(map[string]string)

	// Add key-value pairs to the map
	langs["js"] = "javascript"
	langs["rb"] = "ruby"
	langs["py"] = "python"
	langs["c"] = "c lang"

	// Print the entire map
	fmt.Println("List of languages:", langs)

	// Accessing a value using a key
	fmt.Println("JS stands for:", langs["js"])

	// Delete a key-value pair from the map
	delete(langs, "js") // Removes the entry with key "js"
	fmt.Println("After deleting 'js':", langs)

	// Loop through the map
	// Tip: Maps in Go are unordered — the iteration order is random
	for key, value := range langs {
		fmt.Printf("For key %q, value is %q\n", key, value)
	}
}
