/*
Author: Giovanni De Franceschi
*/












package main

import "fmt"

func main(){
	var age int8 = 127		//max int8 value
	fmt.Println(age)
	age = age + 1
	fmt.Println(age)		//-128

	//age = 999
	//fmt.Println(age)		//error overflow
}