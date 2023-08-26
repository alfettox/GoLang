/*
Author: Giovanni De Franceschi
*/








package main

import (
	"fmt"
)

func main() {
	fmt.Println("Do you own any cats? (y/n)")
	var answer string
	fmt.Scan(&answer)
	fmt.Println("Do you own any dogs? (y/n)")
	var answer2 string
	fmt.Scan(&answer2)

	if answer == "y" && answer2 == "y" {
		fmt.Println("You are a cat and dog person!")
	} else if answer == "y" && answer2 == "n" {
		fmt.Println("You are a cat person!")
	} else if answer == "n" && answer2 == "y" {
		fmt.Println("You are a dog person!")
	} else {
		fmt.Println("You don't own any pets.")
	}
}
