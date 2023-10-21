/*
Author: Giovanni De Franceschi
*/















package main;

import(
	"fmt"

)


func main(){
	var variable1 int;
	fmt.Println("Enter a number. Check if it's 10 or not")
	fmt.Scan(&variable1)

	if (variable1 == 10){
		fmt.Println("variable1 is 10")
	} else {
		fmt.Println("variable1 (", variable1,") is not 10")
	}
}