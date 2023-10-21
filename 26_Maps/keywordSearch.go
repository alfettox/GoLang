/*
Author: Giovanni De Franceschi
*/







package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	myMap := make(map[string]int)
	myMap["a"] = 0
	myMap["b"] = 1
	myMap["c"] = 2
	myMap["d"] = 3
	myMap["e"] = 4
	myMap["f"] = 5
	myMap["g"] = 6
	myMap["h"] = 7
	myMap["i"] = 8
	myMap["j"] = 9

	fmt.Println("Enter a letter to be searched or enter 'q' to quit:")
	var input string
	fmt.Scanln(&input)

	if input == "q" {
		os.Exit(0)
	}

	found := false

	for key, value := range myMap {
		if strings.Contains(key, input) || strings.Contains(fmt.Sprintf("%d", value), input) {
			fmt.Println("Key:", key, "Value:", value)
			found = true
		}
	}

	if found {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

	fmt.Println("Do you want to search again? (y/n)")
	var response string
	fmt.Scanln(&response)

	if response == "y" {
		main()
	} else {
		os.Exit(0)
	}
}
