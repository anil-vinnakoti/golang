package main

import (
	"encoding/json"
	"fmt"
)

// person struct represents a person with name, age, and optional details
type person struct {
	Name    string   `json:"name"`              // Tip: json tag renames the field in the output JSON
	Age     int      `json:"-"`                 // Tip: "-" means this field will be ignored in JSON output
	Details []string `json:"details,omitempty"` // Tip: "omitempty" omits the field if it's empty or nil
}

func main() {
	encodeJson()
}

// encodeJson marshals a slice of person structs into pretty JSON
func encodeJson() {
	persons := []person{
		{Name: "Anil", Age: 27, Details: []string{"dummy", "strings"}},
		{Name: "John", Age: 35, Details: []string{"dummy", "strings"}},
		{Name: "Doe", Age: 29, Details: nil}, // Tip: Details will be omitted due to "omitempty"
	}

	// MarshalIndent formats the JSON with indentation for better readability
	jsonData, err := json.MarshalIndent(persons, "", "\t") // Tip: "" is prefix, "\t" is indentation
	handleError(err)

	// Output the formatted JSON
	fmt.Printf("%s\n", jsonData)
}

// handleError panics if error is not nil
// Tip: Useful for quick error handling in simple programs
func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
