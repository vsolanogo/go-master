package main

import "fmt"

// Program to illustrate how to
// create the slices from array

func main() {
	// Array Creation
	arry := [4]string{"Hello", "from", "Developer", "HFD"}
	// Creating slices from given array
	var myslice1 = arry[1:2]
	myslice2 := arry[0:]
	myslice3 := arry[:2]
	myslice4 := arry[:]
	// Display the result
	fmt.Println("My Array: ", arry)
	fmt.Println("My Slice 1: ", myslice1)
	fmt.Println("My Slice 2: ", myslice2)
	fmt.Println("My Slice 3: ", myslice3)
	fmt.Println("My Slice 4: ", myslice4)
}
