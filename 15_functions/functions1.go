/*
Author: Giovanni De Franceschi
*/





package main

import (
	"fmt"
)


func main (){

	fmt.Println("sum: ", add(3,5))
	fmt.Println("sub: ", sub(3,5))

	fmt.Println(swap("hello", "world"))

	c, d := swap("uno", "due")
	fmt.Println(c, d)
}

func add(x int, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}


func swap (x, y string) (string, string) {
	return y, x
}	