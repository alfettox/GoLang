/*
Author: Giovanni De Franceschi
*/









package main

import (
	"fmt"
	"strconv"
)

func main() {
	var powerOfTwo int64 = 1
	var a int64 = 1

	for a <= 63{
		fmt.Println("2 to the power of " + strconv.FormatInt(a, 10) + " is equal to " + strconv.FormatInt(powerOfTwo, 10))
		powerOfTwo *= 2
		a++
	}

	//2 to the power of 64 overflows int64
}
