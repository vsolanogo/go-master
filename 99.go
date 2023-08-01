package main

import "fmt"

// program to illustrate
// the iterating over a slice using for loop

func main99() {
	// Creating slice
	myslice := []string{"This", "is", "the", "example", "of", "Go", "language"}
	// Iterate using the for loop
	for x := 0; x < len(myslice); x++ {
		fmt.Println(myslice[x])
	}
}
