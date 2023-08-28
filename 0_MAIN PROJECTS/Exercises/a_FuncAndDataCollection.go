/*
Author: Giovanni De Franceschi
*/


package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Enter the limit number to generate a random number:")
	var limit int
	fmt.Scan(&limit)
	randomNumber := randomInt(2 * limit) - limit
	fmt.Println("Random number is:", randomNumber)

	array := createArray()

	fmt.Println("Array:", array)
	sortedArray := sortArray(array)
	fmt.Println("Sorted array:", sortedArray)

	fmt.Println("Max value:", sortedArray[len(sortedArray)-1], ", index:", len(sortedArray)-1)
	fmt.Println("Min value:", sortedArray[0], ", index:", 0)

	reversedArray := reverseArray(sortedArray)
	fmt.Println("Reversed array:", reversedArray)

	fmt.Println("Mean:", mean(array))
	fmt.Println("Median:", median(array))

	uniqueArray := removeDuplicates(array)
	fmt.Println("Array with duplicates removed:", uniqueArray)
}

func randomInt(maxValue int) int {
	return rand.Intn(maxValue)
}

func createArray() []int {
	var array []int
	for i := 0; i < 10; i++ {
		array = append(array, randomInt(100))
	}
	return array
}

func sortArray(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i] = array[i] ^ array[j] // XOR, swap without using a temp variable
				array[j] = array[i] ^ array[j]
				array[i] = array[i] ^ array[j]
			}
		}
	}
	return array
}

func reverseArray(array []int) []int {
	newArray := make([]int, len(array))
	for i := 0; i < len(array); i++ {
		newArray[i] = array[len(array)-1-i]
	}
	return newArray
}

func mean(array []int) float64 {
	sum := 0
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	return float64(sum) / float64(len(array))
}

func median(array []int) float64 {
	if len(array)%2 == 0 {
		return float64(array[len(array)/2]+array[len(array)/2-1]) / 2
	} else {
		return float64(array[len(array)/2])
	}
}

func contains(slice []int, num int) bool {
	for _, v := range slice {
		if v == num {
			return true
		}
	}
	return false
}



func removeDuplicates(array []int) []int {
	var uniqueArray []int

	for _, num := range array {
		if !contains(uniqueArray, num) {
			uniqueArray = append(uniqueArray, num)
		}
	}

	return uniqueArray
}
