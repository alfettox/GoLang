/*
Author: Giovanni De Franceschi
*/








package main

import (
	"fmt"
	"strconv"
)


func main(){
	var input1 int;
	fmt.Println("Enter a number. Check if it's odd or even")

	fmt.Scan(&input1)

	var input1Int = strconv.Itoa(input1)

	if (input1 % 2 == 0){
		fmt.Println(input1Int, "is even")
	} else {
		fmt.Println(input1Int, "is odd")
	}
}