// Program which illustrates
// how to create own panic
// Using the panic function

package main

import "fmt"

// Function
func entry(lang *string, aname *string) {
	// When value of lang
	// is nil it will panic
	if lang == nil {
		panic("Error : The Language cannot be nil")
	}

	// When value of aname
	// is nil it will panic

	if aname == nil {
		panic("Error: The Author name cannot be nil")
	}

	// When values of the lang and aname
	// are non-nil values it will print
	// the normal output

	fmt.Print("THe Author Language: %s \n Author Name: %s\n", *lang, *aname)
}

func main() {
	A_lang := "GO-Language"
	//Here in the entry function, we pass
	// a non-nil, nill values
	// Due to nil value this method panics
	entry(&A_lang, nil)
}
