// program to illustrate the concept of anonymous function recursion

package main

import (
	"fmt"
)

// main function

func main() {
	// declaring anonymous function that takes integer value
	var anon_func func(int)

	// defining the anonymous function that prints the numbers from x to 1
	anon_func = func(number int) {
		// the base case
		if number == 0 {
			return
		} else {
			fmt.Println(number)
			// calling anonymous function recursively
			anon_func(number - 1)
		}
	}

	// call to anonymoys recursice function
	anon_func(5)
}
