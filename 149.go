package main

// Program to show compiler
// throws an error if variable is
// declared but not used

import "fmt"

// the main function
func main() {
	// calling function
	// function returns two values which are
	// assigned to mul and div the identifier
	mul, div := mul_div(110, 9)
	// only using the mul variable
	// compiler will give an error
	fmt.Println("110 x 9 = ", mul)
}

// function returning the two
// values of integer type
func mul_div(nm1 int, nm2 int) (int, int) {
	// returning values
	return nm1 * nm2, nm1 / nm2
}
