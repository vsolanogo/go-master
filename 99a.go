package main

import "fmt"

// Program to illustrate the iterating
// over a slice using range in for loop

func main99a() {
	// Creating a slice
	myslice := []string{"This", "is", "the", "example", "of", "Go", "programing"}
	// Iterate the slice using range in for loop
	for index, ele := range myslice {
		fmt.Printf("Index = %d and element = %s\n",
			index+3, ele)
	}
}
