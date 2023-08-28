/*
Author: Giovanni De Franceschi
*/



package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Declare an uninitialized slice named 's' of int type.
	var s []int
	fmt.Println(s)
	fmt.Println("Slice type is: ", reflect.TypeOf(s))

	// Append values to the slice.
	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 3)
	fmt.Println(s)

	fmt.Println(s[1])
	

	//Empty slice
	var s1 []int
	fmt.Println(s1)
	s1 = append(s1, 1,3)
	fmt.Println(s1)


	copied := make([]int, len(s1))
	//copy a slice
	s2 := copy(copied, s1)

	fmt.Println(s2)
}
