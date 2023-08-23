/*
Author: Giovanni De Franceschi
*/
package main

import (
	"fmt"
)

func main() {
	var currentLetter rune = 'a' // Use rune type to hold a single character

	for a := 0; a < 26; a++ { // Changed the loop limit to 26 for the full alphabet
		fmt.Printf("%d letter %c\n", a+1, currentLetter) // Print the index and the current letter
		currentLetter++
	}
}
