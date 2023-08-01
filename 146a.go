package main

// Program to pass an anonymous
// function as an argument into
// the other function

import "fmt"

// Passing anonymous function
// as argument
func XYZ(i func(a, b string) string) {
	fmt.Println(i("Hello", "for"))
}
func main() {
	value := func(a, b string) string {
		return a + b + "Hello"
	}
	XYZ(value)
}
