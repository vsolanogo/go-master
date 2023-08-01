package main

// illustrate how to give names to return values

import "fmt"

// myfunc return 2 values of the int type
// here, return value name
// is rectangle & square
func myfunc(x, y int) (rectangle int, square int) {
	rectangle = x * y
	square = x * x
	return
}
func main() {
	// The return values are assigned into the twodifferent variables
	var area1, area2 = myfunc(4, 8)
	// Display the values
	fmt.Printf("Area of the rectangle is: %d", area1)
	fmt.Printf("\nThe Area of the square is: %d", area2)
}
