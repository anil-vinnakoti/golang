package main

import (
	"fmt"
	"sort"
)

func main() {
	// Create a slice of strings with initial values
	var list = []string{"a", "b", "c"}
	fmt.Printf("type of list is %T\n", list) // Output: []string

	// Append an item to the slice (adds "d" to the end)
	list = append(list, "d")
	fmt.Println(list) // Output: [a b c d]

	// Slice the original list starting from index 1 to the end
	list = append(list[1:]) // Now list is [b c d]

	// Further slice to keep only the first two elements (index 0 and 1)
	list = append(list[:2]) // Now list is [b c]
	fmt.Println(list)

	// -------------------------------
	// Integer slice example
	// -------------------------------

	// Create a slice of integers using make with length 4
	// Tip: make([]T, length) creates a slice with zero values
	nums := make([]int, 4)

	// Assign values manually
	nums[0] = 1
	nums[1] = 4
	nums[2] = 3
	nums[3] = 2

	// Trying to assign nums[4] = 5 would panic (index out of range)

	// Append values beyond the initial length
	// Tip: append() automatically resizes the underlying array if needed
	nums = append(nums, 5, 6, 7)
	fmt.Println(nums) // Output: [1 4 3 2 5 6 7]

	// Sort the slice in increasing order
	sort.Ints(nums)
	fmt.Println(nums)                     // Output: [1 2 3 4 5 6 7]
	fmt.Println(sort.IntsAreSorted(nums)) // Output: true

	// -------------------------------
	// Removing element from a slice
	// -------------------------------

	// Create a slice of course names
	var courses = []string{"react.js", "javascript", "python", "ruby"}
	fmt.Println(courses)

	// Remove the element at index 2 ("python")
	// Tip: append(slice[:index], slice[index+1:]...) removes the item
	courses = append(courses[:2], courses[3:]...)
	fmt.Println(courses) // Output: [react.js javascript ruby]
}
