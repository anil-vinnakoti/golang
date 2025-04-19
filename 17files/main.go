package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("welcome to files")

	content := "this is a dummy text."

	file, osErr := os.Create("./myFile.txt")

	checkNilError(osErr)
	length, writeErr := io.WriteString(file, content)

	checkNilError(writeErr)
	fmt.Println("length is:", length)

	defer file.Close()

	readFile("./myFile.txt")

}

func readFile(fileName string) {
	dataByte, osErr := os.ReadFile(fileName)
	if osErr != nil {
		panic(osErr)
	}
	fmt.Println("txt data inside file is:", string(dataByte))

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
