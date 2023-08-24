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
	rand.Seed(time.Now().UnixNano())

	randomNumber := rand.Intn(100) + 1

	var guess int
	fmt.Print(">Guess a number between 1 and 100: ")

	for {
		fmt.Println(">Please input your guess")
		fmt.Scan(&guess)

		if guess > randomNumber {
			fmt.Println("Your guess is too HIGH")
		} else if guess < randomNumber {
			fmt.Println("Your guess is too LOW")
		} else {
			fmt.Println("You WIN! The guess is correct. The number is", randomNumber)
			break
		}
	}
}
