/*
Author: Giovanni De Franceschi
*/
package main

import (
	"fmt"
)

func main() {
	myMap := make(map[string]int)
	myMap["key1"] = 42
	myMap["key2"] = 43
	myMap["key3"] = 44

	fmt.Println(myMap)

	_, aExists := myMap["key1"] // Using two variables to check if the key exists
	_, bExists := myMap["key2"] // Using two variables to check if the key exists

	if aExists {
		fmt.Println("a Key exists")
	} else {
		fmt.Println("a Key does not exist")
	}

	if bExists {
		fmt.Println("b Key exists")
	} else {
		fmt.Println("b Key does not exist")
	}
}
