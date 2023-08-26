/*
Author: Giovanni De Franceschi
*/




package main

import (
	"fmt"
)

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	fmt.Println("> ", num)
	return num * factorial(num - 1)
}

func main() {
	num := 4
	fmt.Println("Factorial of", num, "is:")
	fmt.Println("FACTORIAL: ", factorial(num))
}
