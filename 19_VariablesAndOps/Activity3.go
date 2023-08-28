/*
Author: Giovanni De Franceschi
*/













package main

import (
	"fmt"
	"math"
)



func main() {
	var deposit, interestRate, years float64
	fmt.Println("Please insert the initial amount")
	fmt.Scan(&deposit)
	fmt.Println("Please insert the interest rate %")
	fmt.Scan(&interestRate)
	fmt.Println("Please insert the number of years")
	fmt.Scan(&years)
	interestRateFraction := interestRate / 100
	finalAmount := deposit * math.Pow(1+interestRateFraction, years)

	fmt.Printf("The final amount is %.2f\n", finalAmount)
}
