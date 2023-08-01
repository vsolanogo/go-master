package main

import "fmt"

// Program to illustrate how to create slices
// Using the make function

func main() {
	// Creating array of size 7 and slice this array till 4
	// and return the reference of the slice
	// Using make function
	var myslice1 = make([]int, 4, 7)
	fmt.Printf("Slice 1 = %v, \nlength = %d, \ncapacity = %d\n", myslice1, len(myslice1), cap(myslice1))
	// Creating the another array of size 7
	// and return the reference of slice
	// Using the make function
	var myslice2 = make([]int, 7)
	fmt.Printf("Slice 2 = %v, \nlength = %d, \ncapacity = %d\n", myslice2, len(myslice2), cap(myslice2))
}
