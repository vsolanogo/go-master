package main

import (
	"bytes"
	"fmt"
)

// Program to demonstrate the concept of trim in the slice of bytes

func main() {
	// Creating, trimming the slice of bytes
	// Using the Trim function
	rest1 := bytes.Trim([]byte("****Welcome to GoWorld****"), "*")
	rest2 := bytes.Trim([]byte("!!!!Learning how to trim slice of bytes@@@@"), "!@")
	rest3 := bytes.Trim([]byte("^^hello&&"), "$")
	// Display results
	fmt.Printf("Final Slice:\n")
	fmt.Printf("\nSlice 1: %s", rest1)
	fmt.Printf("\nSlice 2: %s", rest2)
	fmt.Printf("\nSlice 3: %s", rest3)
}
