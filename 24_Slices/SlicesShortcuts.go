/*
Author: Giovanni De Franceschi
*/

package main

import ("fmt")

func main() {
	numbers := [] int {0,1,2,3,4,5,6,7,8,9}

	slice1 :=numbers[2:4]				// from index 0 to 4 excluded
	fmt.Println("Slice1: ", slice1)

	slice2 := numbers[:3]
	fmt.Println("Slice2: ", slice2)

	slice3 := numbers[4:]
	fmt.Println("Slice3: ", slice3)

	slice4 := numbers[:]
	fmt.Println("Slice4: ", slice4)

}