/*
Author: Giovanni De Franceschi
*/




package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println("i =", i)
		if i == 5 {
			fmt.Println("Stop value reached i =", i)
			break
		}
        time.Sleep(1000 * time.Millisecond) 

	}
}
