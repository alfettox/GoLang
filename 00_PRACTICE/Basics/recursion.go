/*
Author: Giovanni De Franceschi
*/
package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	var input string
	fmt.Print("Enter an integer: ")
	fmt.Scan(&input)

	inputInt, _ := strconv.Atoi(input)
	result := factorial(inputInt)

	fmt.Println(result)
}

func factorial(input int) *big.Int {
	if input == 0 {
		return big.NewInt(1)
	}
	return new(big.Int).Mul(big.NewInt(int64(input)), factorial(input-1))
}
