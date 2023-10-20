/*
Author: Giovanni De Franceschi
*/










package main

import (
	"fmt"
)

func main() {
	var numUnits int
	fmt.Print("Chose a number? ")
	fmt.Scan(&numUnits) // Scan the user input, _ number of items scanned

	curr := 0
	counter := 0

	for counter < numUnits {
		if curr%3 == 0 && curr%5 == 0 { //if curr divisible by 3 and 5, remainder 0
			fmt.Println("  fizz buzz")
			counter++
		} else if curr%3 == 0 { //if curr divisible by 3, remainder 0
			fmt.Println(" fizz")
			counter++
		} else if curr%5 == 0 { //if curr divisible by 5, remainder 0
			fmt.Println(" buzz")
			counter++
		} else {
			fmt.Println(curr)
		}

		curr++
	}

	fmt.Println("TRADITION!!")
}
