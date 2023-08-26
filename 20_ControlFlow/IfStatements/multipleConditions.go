/*
Author: Giovanni De Franceschi
*/









package main

import (
	"fmt"

)


func main(){
	var input1 int;
	var input2 int;
	fmt.Println("Enter two numbers. Check if they're equal or not")
	fmt.Scan(&input1)
	fmt.Scan(&input2)

	if(input1 == input2){
		fmt.Println(input1, "is equal to", input2)
	} else {
		fmt.Println(input1, "is not equal to", input2)
	}

	if(input1 > input2){
		fmt.Println(input1, "is greater than", input2)
	} else {
		fmt.Println(input1, "is not greater than", input2)
	}

	if(input1 < input2){
		fmt.Println(input1, "is less than", input2)
	} else {
		fmt.Println(input1, "is not less than", input2)
	}

	

}