/*
Author: Giovanni De Franceschi
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var power2 int64 = 1
	var a int64 = 1

	for {
		if a > 10 {
			break
		}

		fmt.Println("2 to the power of " + strconv.FormatInt(a, 10) + " is equal to " + strconv.FormatInt(power2, 10))
		power2 += power2

		a++
	}

	for a := 0; a < 10; a++ {
		if a%2 != 0 {
			continue // Skip odd numbers
		}
		fmt.Println(strconv.Itoa(a) + " is an even number")
	}

	var o int = 20
	var p int = 90

	fmt.Println("o = " + strconv.Itoa(o))
	fmt.Println("p = " + strconv.Itoa(p))

	if o > p {
		goto MESSAGE1
	} else {
		goto MESSAGE2
	}

MESSAGE1:
	fmt.Println("a is greater than b")
MESSAGE2:
	fmt.Println("a is smaller than b")

}
