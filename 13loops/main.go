package main

import "fmt"

func main() {
	// Create a slice of days in a week
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	// Print the full slice
	fmt.Println(days)

	// -----------------------------------------
	// Traditional for loop using index
	// -----------------------------------------
	// Tip: Use this when you need index access with more control
	for d := 0; d < len(days); d++ {
		fmt.Println("the day is", days[d])
	}

	// -----------------------------------------
	// Range-based loop with only index
	// -----------------------------------------
	// Tip: range over slices/arrays returns index by default
	for d := range days {
		fmt.Println("the day is", days[d])
	}

	// -----------------------------------------
	// Range-based loop with index and value
	// -----------------------------------------
	// Tip: Use the comma-ok syntax to access both index and value
	for index, day := range days {
		fmt.Printf("the index is %v and day is %v\n", index, day)
	}

	// -----------------------------------------
	// While-like for loop using only condition
	// -----------------------------------------
	randomValue := 2
	for randomValue < 10 {

		// Tip: Use 'goto' sparingly. It jumps to a label in the code.
		if randomValue == 4 {
			goto ak
		}

		// Tip: 'continue' skips to the next iteration of the loop
		if randomValue == 7 {
			randomValue++
			continue
		}

		// Tip: 'break' exits the loop immediately
		if randomValue == 9 {
			break
		}

		fmt.Println("value is:", randomValue)
		randomValue++
	}

	// Label to jump to using 'goto'
ak:
	fmt.Println("Jumping at AK")
}
