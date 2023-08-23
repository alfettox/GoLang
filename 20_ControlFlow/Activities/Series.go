/*
Author: Giovanni De Franceschi
*/


package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Enter a value for n: ")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	result := 0
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
		fmt.Printf("i = %d, ", i)
		fmt.Printf("sum = %d, ", sum)
		fmt.Printf("result = %d\n", result)
		result += sum
	}

	fmt.Println("Result:", result)
}
