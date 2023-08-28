/*
Author: Giovanni De Franceschi
*/


package main

import (
	"fmt"
)

func main() {
	var word1 string
	var word2 string
	var word3 string
	var word4 string
	var word5 string

	fmt.Println("Enter 5 words:")
	fmt.Scan(&word1, &word2, &word3, &word4, &word5)

	wordSlice := []string{word1, word2, word3, word4, word5}

	totalLength := 0
	for _, word := range wordSlice { // _ is the index, word is the value
		totalLength += len(word) //total length of all words
	}
	average := totalLength / len(wordSlice) //average length of all words

	longWords := make([]string, 0)
	shortWords := make([]string, 0)

	for _, x := range wordSlice {
		if len(x) > average {
			longWords = append(longWords, x)
		} else {
			shortWords = append(shortWords, x)
		}
	}

	fmt.Println("Entered words: ", wordSlice)
	fmt.Println("Average word length: ", average)
	fmt.Println("Long words: ", longWords)
	fmt.Println("Short words: ", shortWords)
}
