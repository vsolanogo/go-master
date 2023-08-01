package main

// Program to illustrate
// the concept of the defer statement

import "fmt"

// Functions
func mul(x1, x2 int) int {
	rest := x1 * x2
	fmt.Println("Result: ", rest)
	return 0
}
func show() {
	fmt.Println("Hello, Everyone")
}

// the main function
func main() {
	// Calling the mul() function
	// Here the mul function behaves
	// like normal function
	mul(43, 25)
	// Calling the mul()function
	// Using defer keyword
	// Here mul() function
	// is defer function
	defer mul(27, 46)
	// Calling show() function
	show()
}
