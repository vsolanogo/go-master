// Program to illustrate the concept of tail recursion
package main

import (
	"fmt"
)

// the tail recursive function to print all the numbers
// from x to 1

func print_num(x int) {
	// if number is still prositive, print it and call the function with decremented value
	if x > 0 {
		fmt.Println(x)
		// last statement in the recursive function
		// tail recursive call
		print_num(x - 1)
	}
}

// the main function

func main() {
	// passing positive nuber, prints 5 to 1
	print_num(5)
}
