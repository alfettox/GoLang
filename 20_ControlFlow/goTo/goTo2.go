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
			goto EVEN
		} else {
			goto ODD
		}

	EVEN:
		fmt.Println(strconv.Itoa(a) + " is even")

	ODD:
		fmt.Println(strconv.Itoa(a) + " is odd")
	}

	fmt.Println("-- End of program --")
}
