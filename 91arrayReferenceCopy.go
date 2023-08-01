package main

// Program to illustrate how to
// copy array by reference

import "fmt"

func main91a() {
	// Creating, initializing an array
	// Using the shorthand declaration
	my_arry1 := [6]int{14, 416, 47, 69, 44, 32}
	// Copying array into new variable
	// Here, elements are passed by reference
	my_arry2 := &my_arry1
	fmt.Println("Array_1: ", my_arry1)

	fmt.Println("Array_2:", *my_arry2)
	my_arry1[5] = 200000
	// when we copy an array
	// into the another array by reference
	// then changes made in original
	// array will reflect in
	// the copy of that array
	fmt.Println("\nArray_1: ", my_arry1)
	fmt.Println("Array_2:", *my_arry2)
}
