/*
Author: Giovanni De Franceschi
*/












package main

import (
	"fmt"
)

func main() {
	var currentLetter rune = 'a'

	for a := 0; a < 26; a++ { 
		fmt.Printf("%d letter %c\n", a+1, currentLetter) 
		currentLetter++
	}
}
