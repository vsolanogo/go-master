package main

// Program to illustrate
// the concept of main() function
// Declaration of main package

// Importing packages
import (
	"fmt"
	"sort"
	"strings"
	"time"
)

// Main function
func main() {
	// Sorting the given slice
	st := []int{335, 79, 113, 14, 86, 12, 467, 9}
	sort.Ints(st)

	fmt.Println("Sorted slice: ", st)
	// Finding the index
	fmt.Println("Index value: ", strings.
		Index("Hello", "ks"))
	// Finding the time
	fmt.Println("Time: ", time.Now().Unix())
}
