/*
Author: Giovanni De Franceschi
*/






package main


import (
	"fmt"
)


func main (){
	var bool1 bool = true
	var bool2 bool = false
	fmt.Println(bool1,bool2)

	fmt.Println(bool1 && bool2)
	fmt.Println(bool1 || bool2)
	fmt.Println(!bool1)
	fmt.Println(!bool2)
	
}