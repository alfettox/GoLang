/*
Author: Giovanni De Franceschi
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	for a := 0; a < 10; a++ {
		if a%2 == 0 {
			fmt.Println(strconv.Itoa(a) + " is even")
		}
	}
}
