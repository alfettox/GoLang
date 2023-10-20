/*
Author: Giovanni De Franceschi
*/



package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args // os.Args is a slice that holds command-line arguments
	programName := args[0]     // The first element is the program name itself
	arguments := args[1:]       // The rest of the elements are the arguments passed

	fmt.Println("Program Name:", programName)
	fmt.Println("Arguments:", arguments)
}
