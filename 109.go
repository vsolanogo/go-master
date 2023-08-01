package main

import (
	"bytes"
	"fmt"
)

// Program to illustrate the concept
// of splitting a slice of bytes

func main() {
	// Creating, initializing the slice of bytes
	// Using the shorthand declaration
	slice_1 := []byte{'!', '!', 'H', 'e', 'l', 'l', 'o', 'f', 'o', 'r', 'W', 'o', 'r', 'l', 'd', '#', '#'}
	slice_2 := []byte{'G', 'r', 'a', 'p', 'e'}
	slice_3 := []byte{'%', 'h', '%', 'e', '%', 'l', '%', 'l', '%', 'o', '%'}
	// Displaying slices
	fmt.Println("Original Slice:")
	fmt.Printf("Slice 1: %s", slice_1)
	fmt.Printf("\nSlice 2: %s", slice_2)
	fmt.Printf("\nSlice 3: %s", slice_3)
	// Splitting slice of bytes
	// Using the Split function
	rest1 := bytes.Split(slice_1, []byte("eek"))
	rest2 := bytes.Split(slice_2, []byte(""))
	rest3 := bytes.Split(slice_3, []byte("%"))
	// Display results
	fmt.Printf("\n\nAfter splitting:")
	fmt.Printf("\nSlice 1: %s", rest1)
	fmt.Printf("\nSlice 2: %s", rest2)
	fmt.Printf("\nSlice 3: %s", rest3)
}
