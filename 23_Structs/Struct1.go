/*
Author: Giovanni De Franceschi
*/



package main

import "fmt"




func main(){

	type structureName struct {
		field1 string
		field2 int
		field3 float64
	}

	var a structureName
	a.field1 = "Hello"
	a.field2 = 5
	a.field3 = 3.14159


	fmt.Println(a);
}