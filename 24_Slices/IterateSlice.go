/*
Author: Giovanni De Franceschi
*/



package main

import(
    "fmt"
)

func main() {
	numbers := [7]int{0,1,2,5,798,43,78}
	fmt.Println("numbers = ", numbers)

	// Slice
	s := numbers[0:4]
	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] = %d\n", i, s[i]) // s[i] is the syntax to access elements of a slice %d is the format specifier for an integer
	}
}

