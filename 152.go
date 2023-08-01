// Program to illustrate
// the multiple defer statements, to illustrate
// LIFO policy

package main

import "fmt"

// Functionals
func add(x1, x2 int) int {
	rest := x1 + x2
	fmt.Println("Result: ", rest)
	return 0
}

// the main function

func main() {
	fmt.Println("Starting")
	// Multiple defer statements
	// Executes in the LIFO order
	defer fmt.Println("Ending")
	defer add(37, 59)
	defer add(12, 12)
}
