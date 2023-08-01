package main

// Program to illustrate
// the concept of call by value

import "fmt"

// function which swap the values
func swap(x, y int) int {
	var o int
	o = x
	x = y
	y = o
	return o
}

// the main function
func main() {
	var a int = 20
	var b int = 30
	fmt.Printf("a = %d and b = %d", a, b)
	// call by values
	swap(a, b)
	fmt.Printf("\n a = %d and b = %d", a, b)
}
