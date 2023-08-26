/*
Author: Giovanni De Franceschi
*/


package main

import "fmt"

func main (){

	var a int = 20 		// a is an integer variable
	b := &a             	// b is a pointer to an integer variable
 
	fmt.Println(b)     	// b is the address of a
	fmt.Println(*b)  // *b is the value stored at the address b refers to
 }
