/*
Author: Giovanni De Franceschi
*/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	l := 0
	for l < 10 {
		var a string
		a = passgenerator() 
		fmt.Println(a)
		l++
	}
}

func passgenerator() string {
	length := 10
	pwd := ""
	for i := 0; i < length; i++ {
		randomChar := string(rand.Intn(26) + 65)
		pwd += randomChar
	}
	return pwd
}
