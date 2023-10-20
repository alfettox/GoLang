/*
Author: Giovanni De Franceschi
*/






package main

import (
	"fmt"
)

func main() {
	var myMap map[string]int
	fmt.Println(myMap) 

	myMap = make(map[string]int) 


	//Assigning values to a map
	myMap["key1"] = 42 
	myMap["key2"] = 43 
	myMap["key3"] = 44 

	fmt.Println(myMap) 

	keys := make([]string, 0, len(myMap))
	for k := range myMap {
		keys = append(keys, k)
	}
	fmt.Println(keys)


	newMap := make(map[string]int) //use make to create a map
	fmt.Println(newMap)

	myMap["key1"] = 42
	myMap["key2"] = 43

	fmt.Print("Updated map: ", myMap, "\n")


	fmt.Println("Length is: ", len(myMap)) 


	for key, value := range myMap	{
		fmt.Println("Key:", key, "Value:", value)
	}
}