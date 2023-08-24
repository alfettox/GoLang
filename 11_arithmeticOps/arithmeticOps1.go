/*
Author: Giovanni De Franceschi
*/












package main

import (
	"fmt"
)

func main() {
	var a, b, c int32 = 20, 10, 4

	fmt.Println("a= ", a)
	fmt.Println("b= ", b)
	fmt.Println("c= ", c)

	fmt.Println("a+b= ", a+b)
	fmt.Println("a-b= ", a-b)
	fmt.Println("a*b= ", a*b)
	fmt.Println("a/b= ", a/b)
	fmt.Println("a%b= ", a%b)

	a++
	fmt.Println("a incremented= ", a)
	a--
	fmt.Println("a decremented= ", a)
}
