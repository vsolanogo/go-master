package main

import "fmt"

func main() {
	// Creating zero value slice
	var myslice []string
	fmt.Printf("Length is = %d\n", len(myslice))
	fmt.Printf("Capacity is = %d ", cap(myslice))
}
