// Program to illtstrate how
// to create Closure
package main

import "fmt"

func main() {
	// Declaring variable
	HFW := 0
	// Assigning an anonymous
	// function to variable
	counter := func() int {
		HFW += 1
		return HFW
	}

	fmt.Println(counter())
	fmt.Println(counter())
}
