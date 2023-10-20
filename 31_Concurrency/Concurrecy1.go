/*
Author: Giovanni De Franceschi
*/

package main

import (
	"fmt"
	"time"
)




func main(){
	go goRoutine()

	fmt.Println("main go-routine")
}


func goRoutine(){
	time.Sleep(2 * time.Second)
	fmt.Println("seondary go-routine")
}