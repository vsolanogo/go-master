package main

// the given string

import (
	"fmt"
	"strings"
)

// the Main method
func main() {
	// Creating, initializing string
	// Using the shorthand declaration
	stry1 := "Welcome, Everyone"
	stry2 := "This is the, example of Golang"
	// Displaying the strings
	fmt.Println("Strings before the trimming:")
	fmt.Println("String 1: ", stry1)
	fmt.Println("String 2: ", stry2)
	// Trimming the prefix string from given strings
	// Using the TrimPrefix() function
	rest1 := strings.TrimPrefix(stry1, "Hello")
	rest2 := strings.TrimPrefix(stry2, "World")
	// Displaying results
	fmt.Println("\nStrings after the trimming:")
	fmt.Println("Result 1: ", rest1)
	fmt.Println("Result 2: ", rest2)
}
