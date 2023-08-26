/*
Author: Giovanni De Franceschi
*/

package main

import "fmt"

func main() {
	// Two-dimensional array
	var twoDimArray [3][3]int

	for i := 0; i < len(twoDimArray); i++ {
		for j := 0; j < len(twoDimArray[i]); j++ {
			twoDimArray[i][j] = (i+1) * (j+1)
		}
	}

	// Print the 2D array in a tabular format
	for i := 0; i < len(twoDimArray); i++ {
		for j := 0; j < len(twoDimArray[i]); j++ {
			fmt.Printf("%d\t", twoDimArray[i][j])
		}
		fmt.Println()
	}
}
