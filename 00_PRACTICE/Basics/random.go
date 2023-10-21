/*
Author: Giovanni De Franceschi
*/


package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(rand.Float32())
	fmt.Println(rand.Intn(10)) //values between 0 and 9 included


	ns := rand.NewSource(time.Now().UnixNano());	//create seed
	generator := rand.New(ns)						//use seed

	fmt.Println(generator.Intn(100))		
}
