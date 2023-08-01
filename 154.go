// Program which illustrates the
// concept of panic

package main

import "fmt"

//the main function

func main() {
	// Creating array of string type
	// Using the var keyboard
	var myarr [3]string
	// Elements are assigned using an index
	myarr[0] = "HE"
	myarr[1] = "Helloeveryone"
	myarr[2] = "Hello"
	// Accessing elements
	// of the array
	// Using the index value

	fmt.Println("The Elements of Array:")
	fmt.Println("The Element 1: ", myarr[0])
	// Program panics because the
	// size of the array is 3
	// we try to access
	// the index 5 which is not
	// available in current array,
	// it gives an runtime error
	fmt.Println("The Element 2: ", myarr[5])
}
