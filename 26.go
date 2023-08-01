package main

import (
	"log"
	"os"
)

// Program to illustrate how to rename,
// move a file in the default directory

func main2() {
	// Rename and Remove a file
	// Using Rename() function
	OriginalPath := "helloo.txt"
	NewPath := "abc.txt"
	es := os.Rename(OriginalPath, NewPath)
	if es != nil {
		log.Fatal(es)
	}
}

func main2() {
	// Rename and Remove file
	// Using Rename() function
	OriginalPath := "/Users/anki/Documents/new_folder/helloo.txt"
	NewPath := "/Users/anki/Documents/new_folder/myfolder/abc.txt"
	es := os.Rename(OriginalPath, NewPath)
	if es != nil {
		log.Fatal(es)
	}
}
