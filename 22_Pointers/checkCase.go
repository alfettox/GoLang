/*
Author: Giovanni De Franceschi
*/




package main

import (
	"fmt"
	"strings"
)

func main() {
	var message string = "HELLO aaRLD"
	messageptr := &message

	// Return true if string is all uppercase
	fmt.Println(isupper(messageptr)) // Pass the pointer to the function

	fmt.Println(isupperStr(message)) // Directly pass the string
}

func isupper(x *string) bool {
	if strings.ToUpper(*x) == *x {
		return true
	}
	return false
}

func isupperStr(x string) bool {
	if strings.ToUpper(x) == x {
		return true
	}
	return false
}
