package main

// Program to check if
// the slice is nil or not

import "fmt"

func main() {
	// creating the slices
	st1 := []int{22, 38, 46}
	var st2 []int
	// If we try to run this commented
	// code compiler will give error
	/*st3:= []int{13, 55, 69}
	  fmt.Println(st1==st3)
	*/
	// Checking if the given slice is nil or not
	fmt.Println(st1 == nil)
	fmt.Println(st2 == nil)
}
