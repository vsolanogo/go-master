package main

// Program to illustrate the concept
// of splitting a slice of bytes

import (
	"bytes"
	"fmt"
)

func main() {
	// Creating, Splitting the slice of bytes
	// Using the Split function
	rest1 := bytes.Split([]byte("****Welcome, to, Tutorial****"),
		[]byte(","))
	rest2 := bytes.Split([]byte("Learning x how x to x trim x a x slice of bytes"),
		[]byte("x"))
	rest3 := bytes.Split([]byte("Helloworld, world"), []byte(""))
	rest4 := bytes.Split([]byte(""), []byte(","))
	// Display results
	fmt.Printf("Final Value:\n")
	fmt.Printf("\nSlice 1: %s", rest1)
	fmt.Printf("\nSlice 2: %s", rest2)
	fmt.Printf("\nSlice 3: %s", rest3)
	fmt.Printf("\nSlice 4: %s", rest4)
}
