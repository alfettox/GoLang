/*
Author: Giovanni De Franceschi
*/





package main


import (
	"fmt"
	"sort"
)

func main(){
	wordsArray := []string {"gatto", "cane", "uccellino", "elefante", "topo", "cavallo", "pecora", "gallina", "maiale"}
	fmt.Println("Original Words: ", wordsArray)

	sort.Strings(wordsArray)
	fmt.Println("Sorted Words in alphabetical order: ", wordsArray)
}