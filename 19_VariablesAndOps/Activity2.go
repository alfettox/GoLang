/*
Author: Giovanni De Franceschi
*/






package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	fmt.Println("Please insert a float number")
	fmt.Scan(&a)
	var integerPart int = int(math.Floor(a))
	var remainder float64 = a - float64(integerPart)
	fmt.Println("The integer part is", integerPart)
	fmt.Println("The remainder is", remainder)
}
