package main

import (
	"fmt"
	"io"
	"os" // Tip: os package handles file system operations
)

func main() {
	fmt.Println("welcome to files")

	// Content to be written to the file
	content := "this is a dummy text."

	// ------------------------------------------
	// Create a new file (or truncate if it exists)
	// Tip: os.Create creates or overwrites the file
	// ------------------------------------------
	file, osErr := os.Create("./myFile.txt")
	checkNilError(osErr)

	// ------------------------------------------
	// Write content to the file using io.WriteString
	// Tip: io.WriteString returns number of bytes written and error
	// ------------------------------------------
	length, writeErr := io.WriteString(file, content)
	checkNilError(writeErr)
	fmt.Println("length is:", length)

	// ------------------------------------------
	// Tip: Always defer file.Close() to release the file resource
	// ------------------------------------------
	defer file.Close()

	// Read and print file content
	readFile("./myFile.txt")
}

// ------------------------------------------
// Read a file and print its content
// Tip: os.ReadFile reads the full file content into memory
// ------------------------------------------
func readFile(fileName string) {
	dataByte, osErr := os.ReadFile(fileName)
	if osErr != nil {
		panic(osErr)
	}
	fmt.Println("txt data inside file is:", string(dataByte))
}

// ------------------------------------------
// Centralized error handling
// Tip: A helper function to panic if error is non-nil
// ------------------------------------------
func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
