package main

// Program to demonstrate how to split a string

import (
	"fmt"
	"strings"
)

// the main function
func main() {
	// Creating, initializing the strings
	stry1 := "Welcome, to the, online session, Helloeveryone"
	stry2 := "My cat name is puffi"
	stry3 := "I like to play chess"
	// Displaying the strings
	fmt.Println("String 1: ", stry1)
	fmt.Println("String 2: ", stry2)
	fmt.Println("String 3: ", stry3)
	// Splitting given strings
	// Using the SplitAfter() function
	rest1 := strings.SplitAfter(stry1, ",")
	rest2 := strings.SplitAfter(stry2, "")
	rest3 := strings.SplitAfter(stry3, "!")
	rest4 := strings.SplitAfter("",
		"Helloeveryone, Hello")
	// Displaying result
	fmt.Println("\nResult 1: ", rest1)
	fmt.Println("Result 2: ", rest2)
	fmt.Println("Result 3: ", rest3)
	fmt.Println("Result 4: ", rest4)
}
