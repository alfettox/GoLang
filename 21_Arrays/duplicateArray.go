/*
Author: Giovanni De Franceschi
*/



package main

import "fmt"

func main() {

	numbers_1 := [3]int{1, 6, 1}
	fmt.Println(numbers_1)

	numbers_2 := numbers_1 	// Copy the values of numbers_1 into numbers_2
	fmt.Println(numbers_2)

	numbers_2[0] = 99

	fmt.Println(numbers_1)
   fmt.Println(numbers_2)
}
