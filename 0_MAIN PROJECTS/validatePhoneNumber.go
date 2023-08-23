/*
Author: Giovanni De Franceschi
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Print("Enter a phone number: ")
	fmt.Scan(&input)

	input = strings.ReplaceAll(input, "-", "") //remove dashes

	if len(input) != 10 { //check length
		goto INVALID
	}

	for i := 0; i < len(input); i++ { 		//check if all digits
		if !isDigit(input[i]) {
			goto INVALID
		}
	}

	fmt.Println("Valid phone number")
	return

INVALID:
	fmt.Println("Invalid phone number")
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}
