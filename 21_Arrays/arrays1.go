/*
Author: Giovanni De Franceschi
*/









package main

import (
	"fmt"
)

func main() {
	var array1 [3]int
	array1[0] = 1
	array1[1] = 2
	array1[2] = 3
	fmt.Println("Original array1:", array1)

	fmt.Println("Length of array1:", len(array1))

	fmt.Println("Element 1:", array1[0])
	fmt.Println("Element 2:", array1[1])
	fmt.Println("Element 3:", array1[2])

	slice1 := array1[:]  // slice1 is a slice of array1,
						 // which is all of array1 (from index 0 to index 2)
						 

	slice1 = append(slice1, 4)

	fmt.Println("Modified slice1:", slice1)
}
