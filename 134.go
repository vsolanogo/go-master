package main

import "fmt"

// Program to illustrate the
// modification concept in map

// the main function
func main() {
	// Creating, initializing a map
	m_a_p := map[int]string{
		90: "Duck",
		91: "Cow",
		92: "Cat",
		93: "Duck",
		94: "Rabbit",
	}
	fmt.Println("Original map: ", m_a_p)
	// Assigned map into a new variable
	new_map := m_a_p
	// Perform the modification in new_map
	new_map[96] = "Monkey"
	new_map[98] = "Donkey"
	// Display after modification
	fmt.Println("The New map: ", new_map)
	fmt.Println("\nModification done in old	map:\n", m_a_p)
}
