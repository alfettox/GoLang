/*
Author: Giovanni De Franceschi
*/
package main

import (
	"fmt"
)

func main() {

	MapFrequence := map[int]string{ // map[keyType]valueType{key1: value1, key2: value2, ...}
		1: "a",
		2: "b",
		3: "c",
		4: "d",
		5: "e",
	}

	fmt.Println(MapFrequence)

	newMap := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	fmt.Println(newMap)

	//same key test
	sameKeysMap := map[string]int{
		"a": 1,
		//"a": 2,  // this will cause an error
		// "a": 3, // this will cause an error
	}

	fmt.Println(sameKeysMap) // the last value is the one that is stored
}
