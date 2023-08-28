/*
Author: Giovanni De Franceschi
*/



package main

import "fmt"



func main(){

	var number int = 10;

	var pointer *int = &number; // & is the address of operator

	fmt.Println("number:", number);
	fmt.Println("pointer_number:", pointer);

	var a int = 20;
	var b *int 			//this *int is a pointer to an integer variable

	fmt.Println(a)
	fmt.Println(b)

	var c int = 30;
	var d *int = &c;
	d = &a; // d is now pointing to a

	fmt.Println(a)
	fmt.Println(c)

}