/*
Author: Giovanni De Franceschi
*/












package main

import(
	"fmt"
)


func main(){
	fmt.Println("Enter a number. I will tell you if it's a valid value or not")
	var input int;
	fmt.Scan(&input)

	if(input > 10){
		fmt.Println("It's a valid positive value")
	} else if (input>= 0){
		fmt.Print("It's a positive value but it's over 10")
	} else {
		fmt.Println("It's a negative value")
	}
}