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
	data := "File content"

	fmt.Println("Original string:", data)

	file, err := os.Create("./myNewFile.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bw := bufio.NewWriter(file)
	n, err := bw.WriteString(data)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bytes written:", n)
	bw.Flush() // Flush the buffered writer

}
