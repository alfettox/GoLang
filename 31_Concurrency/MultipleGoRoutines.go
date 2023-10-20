/*
Author: Giovanni De Franceschi
*/


package main

import (
	"fmt"
	"time"
)

func main() {
	go goRoutine()
	go anotherGoRoutine()

	time.Sleep(3 * time.Second)
	fmt.Println("main go-routine")
}
func goRoutine() {
	fmt.Println("This is my first go-routine")
}
func anotherGoRoutine() {
	fmt.Println("This is my second go-routine")
}
