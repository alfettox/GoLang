/*
Author: Giovanni De Franceschi
*/








package main

import (
	"fmt"
	"strconv"
)

func main(){
	for i := 0; i<100; i++{
		if(i%2 == 0){
			fmt.Println(strconv.Itoa(i) + " is even")
		} else{
			continue
		}
	}
}