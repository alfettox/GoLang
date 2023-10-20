/*
Author: Giovanni De Franceschi
*/


package main

import (
	"fmt"
)

func main(){
	defer fmt.Println("Hello")
	for i:=0; i<10; i++{
		fmt.Println(i)
	}
	defer fmt.Println("World")

}