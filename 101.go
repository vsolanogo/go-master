package main

import "fmt"

func main() {
	// Creating zero value slice
	arry := [6]int{25, 86, 97, 33, 49, 21}
	slc := arry[0:4]
	// Before the modifying
	fmt.Println("Original_Array: ", arry)
	fmt.Println("Original_Slice: ", slc)
	// After the modification
	slc[0] = 10
	slc[1] = 100
	slc[2] = 1000
	fmt.Println("\nNew_Array: ", arry)
	fmt.Println("New_Slice: ", slc)
}
