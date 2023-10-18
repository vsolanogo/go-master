// program to illustrate the concept of head recursion
package main

import (
	"fmt"
)

// the gead recursive function to print all the numbers from 1 to x

func print_num(x int) {
	// if the number is still less than x, call function with decremented value

	if x > 0 {
		// the first statement in function
		print_num(x - 1)
		// printing is done at the returning time
		fmt.Print(x)
	}
}

func main() {
	// pasing positive number, prints 5 to 1
	print_num(5)
}
