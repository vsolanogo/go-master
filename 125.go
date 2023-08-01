// Program to illustrate the concept
// of == and != operator with the strings
package main

import "fmt"

// the main function
func main() {
	// Creating, initializing strings
	// using the shorthand declaration
	str1 := "Hello"
	str2 := "Helo"
	str3 := "Helloeveryone"
	str4 := "Hello"
	// Checking string are equal
	// or not using == operator
	res1 := str1 == str2
	res2 := str2 == str3
	res3 := str3 == str4
	res4 := str1 == str4
	fmt.Println("Result 1: ", res1)
	fmt.Println("Result 2: ", res2)
	fmt.Println("Result 3: ", res3)
	fmt.Println("Result 4: ", res4)
	// Checking the string are not equal
	// using != operator

	res5 := str1 != str2
	res6 := str2 != str3
	res7 := str3 != str4
	res8 := str1 != str4
	fmt.Println("\nResult 5: ", res5)
	fmt.Println("Result 6: ", res6)
	fmt.Println("Result 7: ", res7)
	fmt.Println("Result 8: ", res8)
}
