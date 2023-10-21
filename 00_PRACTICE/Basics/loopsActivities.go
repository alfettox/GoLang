/*
Author: Giovanni De Franceschi
*/

package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		for y := 1; y <= i; y++ {
			fmt.Print(i)
		}
		fmt.Println()
	}


	
}
