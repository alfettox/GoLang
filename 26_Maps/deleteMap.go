/*
Author: Giovanni De Franceschi
*/


package main

import (
	"fmt"
)

func main(){
	myMap := make(map[string]int)
	myMap["key1"] = 42
	myMap["key2"] = 43
	myMap["key3"] = 44
	

	fmt.Println(myMap)


	delete(myMap, "key1")

	fmt.Println(myMap)

}