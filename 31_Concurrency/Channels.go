/*
Author: Giovanni De Franceschi
*/



package main

import (
	"fmt"
)



func main(){
	// var a chan int
	var a = make(chan int)
	fmt.Printf("Channel Type is %T", a)

}