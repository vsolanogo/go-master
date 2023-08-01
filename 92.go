package main

import "fmt"

// Program to illustrate how to pass an
// array as an argument in function

// This function accept
// an array as argument

func myfun(a [5]int, size int) int {
	var x, val, y int
	for x = 0; x < size; x++ {
		val += a[x]
	}
	y = val / size
	return y
}

// the main function
func main() {
	// Creating, initializing an array
	var arr = [5]int{57, 29, 69, 25, 14}
	var rest int
	// Passing an array as an argument
	rest = myfun(arr, 5)
	fmt.Printf("Final result is: %d ", rest)
}
