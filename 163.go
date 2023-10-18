// Program to illustrate
// the concept if indirect recursion

package main

import (
	"fmt"
)

// the recursive function for printing all number upto number x

func print_one(x int) {
	//if number is positive
	// print the number
	// call second function
	if x >= 0 {
		fmt.Println("In first function:", x)
		// call to the second funcion which calls this first function indirectly
		print_two(x - 1)
	}
}

func print_two(x int) {
	// if number is positive print the number, call second function
	if x >= 0 {
		fmt.Println("In second function:", x)
		// call to first function
		print_one(x - 1)
	}
}

func main() {
	//passing positive
	// parameter which prints all the number from 1 - 10
	print_one(10)
	// this will not print anything as it does not fillow base case
	print_one(-1)
}
