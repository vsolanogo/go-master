package main

import "fmt"

// program to illustrate the iterating over
// a slice using a range in for loop without an index

func main100() {
	// Creating slice
	myslice := []string{"This", "is", "the", "example", "of", "Go", "programing"}
	// Iterate the slice
	// using range in for loop without index
	for _, ele := range myslice {
		fmt.Printf("Element = %s\n", ele)
	}
}
