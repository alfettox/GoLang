/*
Author: Giovanni De Franceschi
*/





package main

import (
	"fmt"
	"sort"
)

func main(){
	wordsArray := []string {"cat", "dog", "bird", "elephant", "ant", "lion", "tiger", "zebra", "fish"}
	fmt.Println("Original Words: ", wordsArray)
	fmt.Println("Are sorted? ", sort.StringsAreSorted(wordsArray))

	sort.Strings(wordsArray)
	fmt.Println("Sorted Words: ", wordsArray)
	fmt.Println("Are sorted? ", sort.StringsAreSorted(wordsArray))
}