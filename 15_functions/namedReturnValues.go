/*
Author: Giovanni De Franceschi
*/













package main

import (
	"fmt"
)

// return values can be named
func split (sum int) (x , y int){
	x = sum * 4 / 9
	y = sum - x
	return
}

func main (){

	fmt.Println(split(17))
}