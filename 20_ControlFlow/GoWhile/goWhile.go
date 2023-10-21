/*
Author: Giovanni De Franceschi
*/














package main


import (
	"fmt"
	"time"
)

func main(){
	var a int = 0;

	for a < 10 {
		fmt.Println("a =", a)
		if a == 5 {
			fmt.Println("Stop value reached a =", a)
			break
		}
		time.Sleep(1000 * time.Millisecond) 
		a++
	}
}