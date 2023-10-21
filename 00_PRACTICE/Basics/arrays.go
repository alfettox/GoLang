/*
Author: Giovanni De Franceschi
*/
package main

import ("fmt")

func main(){


	var array1 = []string{"a", "b", "c", "d", "e"}

	for i := 0; i < len(array1); i++ {
		fmt.Println(i)
	}

	fmt.Println(array1[:3])
	fmt.Println(array1[2:])

	newSlice := sliceSteps(array1,2)
	fmt.Println(newSlice)

	newSlice2 := sliceSteps(array1,3)
	fmt.Println(newSlice2)


	var emptyArray[5] int

	fmt.Println(emptyArray)
}

func sliceSteps(source []string, step int) []string {
	var result []string
	for i := 0; i < len(source); i += step {
		result = append(result, source[i])
	}
	return result
}