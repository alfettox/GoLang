/*
Author: Giovanni De Franceschi
*/
package main

import (
	"bufio" 		// Used to read input from the console
	"fmt"
	"os"
)

func main() {
	var input string

	fmt.Print("Enter a string: ")
	scanner := bufio.NewScanner(os.Stdin) 		// Used to read input from the console
	if scanner.Scan() {
		input = scanner.Text()
	}

	for _, char := range input {
		fmt.Println(string(char))
	}
}
