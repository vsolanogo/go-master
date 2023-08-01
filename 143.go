package main

// Program to illustrate
// the concept of variadic function

import (
	"fmt"
	"strings"
)

// Variadic function to join the strings
func joinstr(element ...string) string {
	return strings.Join(element, "-")
}

func main() {
	// zero argument
	fmt.Println(joinstr())
	// the multiple arguments
	fmt.Println(joinstr("Hello", "HEW"))
	fmt.Println(joinstr("Hello", "Everyone", "World"))
	fmt.Println(joinstr("H", "E", "L", "L", "O"))
}
