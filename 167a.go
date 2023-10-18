// program which illustrates the concept of recover
package main

import (
	"fmt"
)

// this function handles panic occur in the entry function
// with help of the recover function
func handlepanic() {
	if x := recover(); x != nil {
		fmt.Println("RECOVER", x)
	}
}

// function

func entry(lang *string, aname *string) {
	// the defferred function
	defer handlepanic()

	// when value of lang is nil it will panic
	if lang == nil {
		panic("Error: The language cannot be nil")
	}

	// when value of aname is nil it will panic
	// is nil it will panic

	if aname == nil {
		panic("Error: The author name cannot be nil")
	}

	fmt.Printf("The author language: %s \n Author Name: %s\n", *lang, *aname)

	fmt.Printf("The Return successfully from entry function")
}

// the main function
func main() {
	A_lang := "GO-Language"
	entry(&A_lang, nil)
	fmt.Prtintf("The Return successfully from main function")
}
