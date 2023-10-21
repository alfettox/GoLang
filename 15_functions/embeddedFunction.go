/*
Author: Giovanni De Franceschi
*/










package main

import (
	"fmt"
	"math/rand"
)

func passwordGenerator() func() string {
	var length int = 8
	return func() string {
		password := ""
		for i := 0; i < length; i++ {
			randomChar := string(rand.Intn(94) + 33) // visibles ASCII characters
			password += randomChar
		}
		return password
	}
}

func main() {
	passwordGenerator := passwordGenerator() // embedded function (closure)
	fmt.Println(passwordGenerator())

}


