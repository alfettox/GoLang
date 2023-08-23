/*
Author: Giovanni De Franceschi
*/







package main

import (
	"fmt"
	"strconv"
)

func main() {
	var firstNumber string
	var secondNumber string

	fmt.Print("Enter the first number: ")
	fmt.Scan(&firstNumber)

	fmt.Print("Enter the second number: ")
	fmt.Scan(&secondNumber)

	fmt.Println(firstNumber + secondNumber) // concatenation

	firstNumberInt, err1 := strconv.Atoi(firstNumber)
	if err1 != nil {
		fmt.Println("Error converting firstNumber:", err1)
		return
	}

	secondNumberInt, err2 := strconv.Atoi(secondNumber)
	if err2 != nil {
		fmt.Println("Error converting secondNumber:", err2)
		return
	}

	fmt.Println(firstNumberInt + secondNumberInt) // addition
}
