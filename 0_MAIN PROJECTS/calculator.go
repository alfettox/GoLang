/*
Author: Giovanni De Franceschi
*/



package main

import (
	"fmt"
	"os"
)

func main() {
	for {
		processInput()
	}
}

func processInput() {
	var num1, num2 int
	var operator string

	// Input for num1
	fmt.Print("Enter the first number: ")
	_, err := fmt.Scan(&num1) // _ is a blank identifier allowing us to ignore the number of items scanned
	if err != nil {
		fmt.Println("Invalid input for the first number.")
		return
	}

	// Input for operator
	fmt.Print("Enter the operator: ")
	_, err = fmt.Scan(&operator)
	if err != nil {
		fmt.Println("Invalid input for the operator.")
		return
	}

	// Input for num2
	fmt.Print("Enter the second number: ")
	_, err = fmt.Scan(&num2)
	if err != nil {
		fmt.Println("Invalid input for the second number.")
		return
	}

	switch operator {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		if num2 == 0 {
			fmt.Println("ERROR: divide by zero.")
			return
		}
		fmt.Println(num1 / num2)
	default:
		fmt.Println("Invalid operator")
		return
	}

	var input string
	fmt.Print("Enter 'quit' or 'exit' to quit. Or any key to continue: ")
	fmt.Scan(&input)

	if input == "quit" || input == "exit" {
		fmt.Println("Closing program.")
		os.Exit(0)
	} else {
		fmt.Println("You entered:", input)
	}
}
