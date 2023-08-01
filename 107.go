package main

import (
	"bytes"
	"fmt"
)

// Program to demonstrate the concept of trim in the slice of bytes

func main() {
	// Creating, initializing
	// the slice of bytes
	// Using the shorthand declaration
	slice_1 := []byte{'!', '!', 'H', 'e', 'e', 'l', 'o', 'o', 'o', 'r', 'W', 'o', 'r', 'l',
		'd', '#', '#'}
	slice_2 := []byte{'*', '*', 'G', 'r', 'a',
		'p', 'e', '^', '^'}
	slice_3 := []byte{'%', 'h', 'e', 'l', 'l',
		'o', '%'}
	// Displaying slices
	fmt.Println("The Original Slice:")
	fmt.Printf("Slice 1: %s", slice_1)
	fmt.Printf("\nSlice 2: %s", slice_2)
	fmt.Printf("\nSlice 3: %s", slice_3)
	// Trimming the specified leading
	// and trailing Unicodes points
	// from given slice of bytes
	// Using Trim function
	rest1 := bytes.Trim(slice_1, "!#")
	rest2 := bytes.Trim(slice_2, "*^")
	rest3 := bytes.Trim(slice_3, "@")
	// Display results
	fmt.Printf("New Slice:\n")
	fmt.Printf("\nSlice 1: %s", rest1)
	fmt.Printf("\nSlice 2: %s", rest2)
	fmt.Printf("\nSlice 3: %s", rest3)
}
