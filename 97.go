package main

// Program to illustrate how to
// create slices from slice

import "fmt"

func main() {
	// Creating s slice
	oRignAl_slice := []int{80, 20, 50, 10, 54, 89, 70}
	// Creating the slices from the given slice
	var myslice1 = oRignAl_slice[1:5]
	myslice2 := oRignAl_slice[0:]
	myslice3 := oRignAl_slice[:6]
	myslice4 := oRignAl_slice[:]
	myslice5 := myslice3[2:4]
	// Display result
	fmt.Println("Original Slice:", oRignAl_slice)
	fmt.Println("New Slice 1:", myslice1)
	fmt.Println("New Slice 2:", myslice2)
	fmt.Println("New Slice 3:", myslice3)
	fmt.Println("New Slice 4:", myslice4)
	fmt.Println("New Slice 5:", myslice5)
}
