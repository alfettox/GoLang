/*
Author: Giovanni De Franceschi
*/



package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./flatland01.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // Defer the file close operation

	b1 := make([]byte, 10)
	data, err := file.Read(b1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf(string(b1[:data]))
}
