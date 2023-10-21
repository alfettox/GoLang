/*
Author: Giovanni De Franceschi
*/














package main

import (
	"fmt"
	"time"
)

func main() {
	var a int = 0

	for ; a < 15; a++ {
		fmt.Print(a)
		fmt.Println() // Flush the buffer to display the number immediately
		time.Sleep(1000 * time.Millisecond)
		if a%2 == 0 {
			fmt.Print(" Even | ")
		} else {
			fmt.Print(" Odd | ")
		}
	}
}
