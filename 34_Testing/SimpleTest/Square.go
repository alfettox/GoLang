/*
Author: Giovanni De Franceschi
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println("main function started")
	fmt.Println(Square(2))
}

func Square(number int) int {
	return number * number
}
