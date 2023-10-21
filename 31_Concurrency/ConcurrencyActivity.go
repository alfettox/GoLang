/*
Author: Giovanni De Franceschi
*/



package main

import (
	"fmt"
	"time"
)

func main() {
	go countDown() 
	go countUp()
	time.Sleep(50 * time.Second)

	fmt.Println("completed")
}

func countDown() {
	for i := 100; i >= 0; i-- {
		fmt.Println("Countdown:", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func countUp() {
	for i := 0; i <= 100; i++ {
		fmt.Println("Countup:", i)
		time.Sleep(110 * time.Millisecond)
	}
}
