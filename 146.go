package main

// Program to pass arguments
// in anonymous function

import "fmt"

func main() {
	// Passing arguments in the anonymous function
	func(ele string) {
		fmt.Println(ele)
	}("Helloeveryone")
}
