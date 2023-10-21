/*
Author: Giovanni De Franceschi
*/








package main

import "fmt"

func main() {
    firstArray := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    var reversedArray [10]int

    for index, value := range firstArray {
        reversedArray[len(reversedArray)-1-index] = value 
		// Reverse the array by assigning the values in reverse order, using the index 
		// -1 to account for the zero-based index of the array
    }

    fmt.Println("Reversed Array:", reversedArray)

	arrayToBeModified := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	arrayToBeModified[4] = 100 // Modify the value at index 4

	fmt.Println("Modified Array:", arrayToBeModified)
}
