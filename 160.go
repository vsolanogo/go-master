// Program to show  how to create recursive Anonymous func

package main

import "fmt"

func main() {
	//anonymous func
	var recursiveAnonymous func()
	recursiveAnonymous = func() {
		// printing message to show the function call and iteration
		fmt.Println("The anonymous functions could be recursive.")
		//calling the same function recursively
		recursiveAnonymous()
	}
	recursiveAnonymous()
}
