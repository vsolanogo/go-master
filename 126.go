package main

// Program to illustrate concept
// of comparison operator with the strings

import "fmt"

// the main function
func main() {
	// Creating, initializing
	// slice of string using
	// the shorthand declaration
	myslice := []string{"Hello", "Hello",
		"hfw", "HFW", "from"}
	fmt.Println("Slice: ", myslice)
	// Using the comparison operator
	result1 := "HFW" > "Hello"
	fmt.Println("Result 1: ", result1)
	result2 := "HFW" < "hello"
	fmt.Println("Result 2: ", result2)
	result3 := "Hello" >= "from"
	fmt.Println("Result 3: ", result3)
	result4 := "Hello" <= "from"
	fmt.Println("Result 4: ", result4)
	result5 := "Hello" == "Hello"
	fmt.Println("Result 5: ", result5)
	result6 := "Hello" != "from"
	fmt.Println("Result 6: ", result6)
}
