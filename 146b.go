package main

// Program to illustrate
// the use of anonymous function

import "fmt"

// Returning the anonymous function
func XYZ() func(a, b string) string {
	myf := func(a, b string) string {
		return a + b + "Everyone"
	}
	return myf
}

func main() {
	value := XYZ()
	fmt.Println(value("Hello ", "to "))
}
