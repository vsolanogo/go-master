// program to show how to create recursive anonymous function
package main

import (
	"fmt"
)

func main() {
	// the anonymous function
	var recursiveAnonymous func(int)
	// passing arguments to anonymous function
	recursiveAnonymous = func(variable int) {
		//checking condition to return
		if variable == -1 {
			fmt.Println("welcome to our channel")
			return
		} else {
			fmt.Println(variable)
			//calling the same functionality recursively
			recursiveAnonymous(variable - 1)
		}
	}

	// the main calling of function
	recursiveAnonymous(10)
}
