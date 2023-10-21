/*
Author: Giovanni De Franceschi
*/







package main

import (
	"fmt"
)

func main() {
	numbers := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(numbers)

	slice1 := numbers[4:6]
	fmt.Println(slice1)

	slice1 = slice1[:cap(slice1)] // re-slice extending till capacity of slice using cap() function
	fmt.Println(slice1)

	/* cap is a built-in function cap() that returns the capacity of a slice. 
	The capacity of a slice is the number of elements in the underlying array,
	counting from the first element in the slice. */
}