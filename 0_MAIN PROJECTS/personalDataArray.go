/*
Author: Giovanni De Franceschi
*/


package main

import (
	"fmt"
	"os"
)

func main() {
	var name string
	fmt.Println("Enter your name: ")
	fmt.Scan(&name)

	var birthdayDate string
	fmt.Println("Enter your birthday date: ")
	fmt.Scan(&birthdayDate)

	var city string
	fmt.Println("Enter your city: ")
	fmt.Scan(&city)

	var quit string
	fmt.Println("Write 'quit' to exit or any key to continue ")
	fmt.Scan(&quit)
	if quit == "quit" {
		os.Exit(0)
	}

	var hobby string
	fmt.Println("Enter your hobby: ")
	fmt.Scan(&hobby)

	answersArray := [...]string{name, birthdayDate, city, hobby}

	fmt.Println("Your name is:", answersArray[0])
	fmt.Println("Your birthday date is:", answersArray[1])
	fmt.Println("Your city is:", answersArray[2])
	fmt.Println("Your hobby is:", answersArray[3])
}
