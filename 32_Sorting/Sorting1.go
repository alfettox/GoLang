/*
Author: Giovanni De Franceschi
*/
package main

import (
	"fmt"
	"sort"
)

func main(){
		numbers := []int {91, 3, 12, 5, 44, 17, 66, 79, 8}

		fmt.Println("Original Numbers: ", numbers)
		sort.Ints(numbers)
		fmt.Println("Sorted Numbers: ", numbers)

}