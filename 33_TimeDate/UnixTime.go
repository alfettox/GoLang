/*
Author: Giovanni De Franceschi
*/
package main

import (
	"fmt"
	"time"
)

func main(){
	now := time.Now()
	unixTime := now.Unix()
	fmt.Println("Current Time Unix: ", unixTime)

	unixTimeNano := now.UnixNano()
	fmt.Println("Current Time Unix Nano:", unixTimeNano)

}