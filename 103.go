package main

// Program to illustrate how to sort
// elements present in the slice

import (
	"fmt"
	"sort"
)

func main() {
	// Creating the Slice
	slc1 := []string{"C++", "Java", " Python ",
		"Go", "Python"}
	slc2 := []int{35, 87, 13, 91, 34, 41, 86, 58, 69}
	fmt.Println("Before the sorting:")
	fmt.Println("Slice 1: ", slc1)
	fmt.Println("Slice 2: ", slc2)
	// Performing sort operation on the
	// slice using the sort function
	sort.Strings(slc1)
	sort.Ints(slc2)
	fmt.Println("\nAfter sorting:")
	fmt.Println("Slice 1: ", slc1)
	fmt.Println("Slice 2: ", slc2)
}
