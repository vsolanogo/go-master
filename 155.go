// Program which illustrates
// the concept of Defer while panicking
package main

import (
	"fmt"
)

//Function

func entry(lang *string, aname *string) {
	// the Defer statement
	defer fmt.Println("The Defer statement in the entry function")
	// When value of lang
	// is nil it will panic
	if lang == nil {
		panic("Error: The Language cannot be nil")
	}
	//When value of aname
	// is nil it will panic
	if aname == nil {
		panic("Error: The Author name cannot be nil")
	}
	// When value sof the lang and aname
	// are non-nil value sit will
	// print the normal output
	fmt.Printf("The Author Language: %s \n Author name:  %s\n", *lang, *aname)
}

//the main function

func main() {
	A_lang := "GO-Language"
	// the Defer statement
	defer fmt.Println("the Defer statement in the main function")
	// in entry function, we pass
	// one non-nil and one-nil value
	// Due to nil value this method panics
	entry(&A_lang, nil)
}
