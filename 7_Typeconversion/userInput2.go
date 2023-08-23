/*
Author: Giovanni De Franceschi
*/







package main

import (
	"fmt"
	"strconv"
)

func main(){
	fmt.Println("Enter the first number: ")
	var firstNumber string
	fmt.Scan(&firstNumber)

	fmt.Println("Enter the second number: ")
	var secondNumber string
	fmt.Scan(&secondNumber)

	var firstNumberInt int
	var secondNumberInt int

	firstNumberInt, _ = strconv.Atoi(firstNumber)
	secondNumberInt, _ = strconv.Atoi(secondNumber)

	var sum int;

	sum = firstNumberInt + secondNumberInt;
	fmt.Println("The sum is: ", sum)
}