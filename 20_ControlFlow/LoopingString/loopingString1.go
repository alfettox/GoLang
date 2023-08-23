/*
Author: Giovanni De Franceschi
*/


package main

import (
	"fmt"
)

func main() {
	var message string = "This is a message"
	fmt.Println(message)

	for i := 0; i < len(message); i++ {
		fmt.Println("Index =", i, "message[i] =", string(message[i]))
	}
}
