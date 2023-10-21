/*
Author: Giovanni De Franceschi
*/
package main

import "fmt"

func main() {
	sumNumbers(2, 5, 4)
	sumNumbers(1, 2, 3, 4, 5)
}

func sumNumbers(numbers ...int) {			//accept variable number of params
	sum := 0
	for i, num := range numbers {
		fmt.Println("Current element:", num, "| index: ", i)
		sum += num
	}

	fmt.Println("Sum of values: ", sum)
}
