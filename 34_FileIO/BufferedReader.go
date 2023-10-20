/*
Author: Giovanni De Franceschi
*/


package main

import (
	"bufio"
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

	br := bufio.NewReader(file)
	data, err := br.Peek(5)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}
