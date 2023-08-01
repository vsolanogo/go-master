package main

// Program to illustrate
// the concept of init() function
// Declaration of main package

// the importing package
import "fmt"

// the multiple init() function
func init() {
	fmt.Println("Welcome everyone")
}
func init() {
	fmt.Println("Hello everyone ")
}

// the main function
func main() {
	fmt.Println("Welcome to home")
}
