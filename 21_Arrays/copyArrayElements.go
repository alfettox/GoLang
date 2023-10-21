/*
Author: Giovanni De Franceschi
*/










package main

import (
	"fmt"
)

func main() {

	var array [5]int

	array[0] = 1
	array[1] = 2
	array[2] = 3
	array[3] = 4
	//element 5 is empty = 0

	fmt.Println("Original array:", array)

	var newArray [7]int

	for i := 0; i < len(array); i++ {
		if array[0] != 0 {
			newArray[i] = array[i]
		}
	}
	newArray[4] = 5
	newArray[5] = 6
	newArray[6] = 7

	fmt.Println("New array:", newArray)

	//replace element
	newArray[2] = 10
	fmt.Println("New array:", newArray)

}
