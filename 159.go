// Program to illustrate how to create the data isolation

package main

import "fmt"

// newCounter function to
// isolate the global variable
func newCounter() func() int {
	HFW := 0
	return func() int {
		HFW += 1
		return HFW
	}
}

func main() {
	// newCOunter function is assigned to a variable
	counter := newCounter()
	// invoke the counter
	fmt.Println(counter())
	//invoke the counter
	fmt.Println(counter())
}
