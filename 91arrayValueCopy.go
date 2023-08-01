package main

import "fmt"

// Program to illustrate how
// to copy array by value

func main91() {
	// Creating, initializing an array
	// Using the shorthand declaration
	my_arry1 := [5]string{"C", "Go", "Java", " Scala ", "C#"}
	// Copying array into new variable
	// Here, elements are passed by value
	my_arry2 := my_arry1
	fmt.Println("Array_1: ", my_arry1)
	fmt.Println("Array_2:", my_arry2)
	my_arry1[0] = "C++"
	// when we copy an array
	// into the another array by value then changes made in the original
	// array do not reflect in copy of that array
	fmt.Println("\nThe Array_1: ", my_arry1)
	fmt.Println("The Array_2: ", my_arry2)
}
