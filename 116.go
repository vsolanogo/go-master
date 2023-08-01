package main

// Program to illustrate
// how to trim string

import (
	"fmt"
	"strings"
)

// the main method
func main() {
	// Creating, initializing string
	// Using the shorthand declaration
	stry1 := "!!Welcome to Everyone !!"
	stry2 := "@@This is the example of Golang$$"
	// Displaying strings
	fmt.Println("Strings before the trimming:")
	fmt.Println("String 1: ", stry1)
	fmt.Println("String 2:", stry2)
	// Trimming given strings
	// Using Trim() function
	rest1 := strings.Trim(stry1, "!")
	rest2 := strings.Trim(stry2, "@$")
	// Displaying results
	fmt.Println("\nStrings after the trimming:")
	fmt.Println("Result 1: ", rest1)
	fmt.Println("Result 2:", rest2)
}
