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
	decodeJson()
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

// decodeJson unmarshals JSON into struct and map forms
func decodeJson() {
	// JSON input simulating data received from the web
	jsonDataFromWeb := []byte(`
		{
			"name": "Anil",
			"details": ["dummy","strings"]
		}
	`)

	var personA person

	// --------------------------------------------
	// Tip: json.Valid checks if JSON is correctly formatted
	// --------------------------------------------
	isJsonValid := json.Valid(jsonDataFromWeb)
	if isJsonValid {
		fmt.Println("JSON was valid ✅")

		// --------------------------------------------
		// Unmarshal JSON into a predefined struct
		// Tip: Use when you know the structure ahead of time
		// --------------------------------------------
		json.Unmarshal(jsonDataFromWeb, &personA)
		fmt.Printf("Struct output: %#v\n", personA)
	} else {
		fmt.Println("JSON was not valid ❌")
	}

	// --------------------------------------------
	// Unmarshal JSON into a map[string]interface{}
	// Tip: Use when structure is unknown or dynamic
	// --------------------------------------------
	var onlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &onlineData)
	fmt.Printf("Map output: %#v\n", onlineData)

	// Loop through map to inspect keys and dynamic values
	for key, value := range onlineData {
		fmt.Printf("Key: %v | Value: %v | Type: %T\n", key, value, value)
	}
}

// handleError panics if error is not nil
// Tip: Useful for quick error handling in simple programs
func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
