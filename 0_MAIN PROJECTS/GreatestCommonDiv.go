/*
Author: Giovanni De Franceschi
*/









package main

import (
	"fmt"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	var number1, number2 int

	fmt.Print("Enter the first number: ")
	fmt.Scan(&number1)
	fmt.Print("Enter the second number: ")
	fmt.Scan(&number2)

	fmt.Println("The greatest common divisor is", gcd(number1, number2))
}


