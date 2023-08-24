/*
Author: Giovanni De Franceschi
*/









package main

import (
	"fmt"
	"math/rand"
)

func main(){
	fmt.Println(rand.Float64()) // values between 0.0 and 1.
	fmt.Println(rand.Int()) // values between 0 and 2^32-1
	fmt.Println(rand.Intn(10)) // 0 <= n < 10
	fmt.Println(rand.Intn(10)+5) // 5 <= n < 15
}