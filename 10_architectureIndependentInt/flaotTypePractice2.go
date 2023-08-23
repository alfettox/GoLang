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

	fmt.Print("Enter the first number:")
	fmt.Scan(&firstNumber)

	fmt.Print("Enter the second number:")
	fmt.Scan(&secondNumber)

	fmt.Println(firstNumber + secondNumber)

	firstNumberInt, err1 := strconv.Atoi(firstNumber)
	if err1 != nil {
		fmt.Println("Error parsing int firstNumber: ", err1)
		return
	}
	secondNumberInt, err2 := strconv.Atoi(secondNumber)
	if err2 != nil {
		fmt.Println("Error parsing int secondNumber: ", err2)
		return
	}

	fmt.Println(firstNumberInt + secondNumberInt)
}
