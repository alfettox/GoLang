/*
Author: Giovanni De Franceschi
*/





package main

import (
	"fmt"
)

func fibonacci(num int) int {
	if num <= 0 {
		return 0
	} else if num == 1 {
		return 1
	}
	return fibonacci(num-1) + fibonacci(num-2)
}

func main() {
	var num int;
	fmt.Println("Recursion: Fibonacci sequence")
	fmt.Println("Enter the number of elements: ")
	fmt.Scanln(&num)
	fmt.Println("Fibonacci sequence up to:", num)
	for i := 0; i <= num; i++ {
		fmt.Println(fibonacci(i))
	}
}
