/*
Author: Giovanni De Franceschi
*/

package main

import (
	"fmt"
	"sort"
)

// Define a custom type named mytype, which is just a slice of strings.
type mytype []string

// Implement the sort.Interface methods (Len, Less, Swap) for the custom type mytype.
func (s mytype) Less(i, j int) bool {
	return len(s[i]) < len(s[j]) // Compare strings by their lengths.
}

func (s mytype) Swap(i, j int) {
	s[i], s[j] = s[j], s[i] // Swap two elements in the slice.
}

func (s mytype) Len() int {
	return len(s) // Return the length of the slice.
}

func main() {
	// Define a slice of strings representing car brands.
	carsArray := []string{"Ford", "BMW", "Audi", "Honda", "Toyota", "Mercedes", "Volkswagen", "Ferrari", "Lamborghini", "Porsche"}

	fmt.Println("Original Cars:", carsArray)

	// Convert the carsArray to the custom type mytype.
	myCars := mytype(carsArray)

	// Use the sort.Sort function to sort the slice based on the custom sorting rules.
	sort.Sort(myCars)

	fmt.Println("Sorted Cars:", myCars)
}
