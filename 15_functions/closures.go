/*
Author: Giovanni De Franceschi
*/









package main

import (
	"fmt"
)

func main() {
	var a int = 5
	var b int = 10

	var add = func() int { // anonymous function (closure)
		return a + b
	}

	fmt.Println(add())
}
