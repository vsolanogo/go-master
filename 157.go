// Program which illustrates
// the concept of recover
package main

import "fmt"

// This function is created to handle panic occurs in entry function
// but it does not handle panic occurred in entry function
// because it called in normal function
func handlepanic() {
	if a := recover(); a != nil {
		fmt.Println("RECOVER", a)
	}
}

// Function
func entry(lang *string, aname *string) {
	// Normal function
	handlepanic()
	// When value of lang
	// is nil it will panic
	if lang == nil {
		panic("Error: Language cannot be nil")
	}
	// When value of aname
	// is nil it will panic
	if aname == nil {
		panic("Error: Author name cannot be nil")
	}
	fmt.Printf("The Author Language: %s \n Author Name: %s\n", *lang, *aname)
	fmt.Printf("Return successfully from entry function")
}

// The main function
func main() {
	A_lang := "GO Language"
	entry(&A_lang, nil)
	fmt.Printf("Return successfully from the main function")
}
