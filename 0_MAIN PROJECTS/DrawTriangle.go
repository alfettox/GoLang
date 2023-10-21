/*
Author: Giovanni De Franceschi
*/













package main

import (
	"fmt"
)


func main(){
	var i int;
	var j int;

	for i=0; i < 10; i++{
		for j=0; j < i; j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}