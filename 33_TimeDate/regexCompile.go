/*
Author: Giovanni De Franceschi
*/




package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Create a regular expression pattern to match a specific format.
	r := regexp.MustCompile(`S\d+|It's five o'clock now:`)
	fmt.Println("Search term:", r)

	fmt.Println("S54366456SDfhdgstf7986:", r.MatchString("S54366456SDfhdgstf7986"))
	fmt.Println("It's five o'clock now:", r.MatchString("It's five o'clock now:"))

	fmt.Println("The phone number is 555-9980", r.FindString("The phone number is 555-9980"))
	fmt.Println("Alexander Hamilton:", r.FindString("Alexander Hamilton"))
}
