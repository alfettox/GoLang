/*
Author: Giovanni De Franceschi
*/





package main

import (
	"fmt"
	"strconv"
)

func main(){
	for a := 0; a <10 ; a++{

		if( a % 2 == 0){
			continue 	//continue statement skips the rest of the loop body and goes to the next iteration
		}
		fmt.Println(strconv.Itoa(a) + " is odd")
	}
}