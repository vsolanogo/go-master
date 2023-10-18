// Program to illustrate the concept of direct recursion
package main

import (
	"fmt"
)

// the recursive function for calculating factorial of a positive integer
func factorial_calc(number int) int {
	// this is base condition if number is 0 or 1 the functiion will return 1
	if number == 0 || number == 1 {
		return 1
	}
	// if the negative argument is given, it prints error message & return - 1
	if number < 0 {
		fmt.Println("Invalid-number")
		return -1
	}

	//  the recursive call to itself with argument decremented
	// by 1 integer so that it eventually reaches base case
	return number * factorial_calc(number-1)
}

func main() {
	// passing 0 as a parameter
	answer1 := factorial_calc(0)
	fmt.Println(answer1, "\n")
	// passing a positive integer
	answer2 := factorial_calc(5)
	fmt.Println(answer2, "\n")
	// passing negative integer prints error message with return value of -1
	answer3 := factorial_calc(-1)
	fmt.Println(answer3, "\n")
	//passing positive integer
	answer4 := factorial_calc(10)
	fmt.Println(answer4, "\n")
}
