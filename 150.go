package main

// Program to the use of Blank identifier

import "fmt"

// the main function
func main() {
	// calling function
	// function returns two values which are
	// assigned to mul and blank identifier
	mul, _ := mul_div(110, 8)
	// only using the mul variable
	fmt.Println("110 x 8 = ", mul)
}

// function returning the two
// values of integer type
func mul_div(nm1 int, nm2 int) (int, int) {
	// returning the values
	return nm1 * nm2, nm1 / nm2
}
