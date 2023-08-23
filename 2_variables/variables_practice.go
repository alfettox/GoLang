/*
Author: Giovanni De Franceschi
*/






package main

import "fmt"

func main() {
	var firstName string = "Giovanni"
	var lastName string = "De Franceschi"
	var birthYear = 1985
	var currentYear = 2023
	var age = currentYear - birthYear

	fmt.Println("My name is: ", firstName, lastName)
	fmt.Println("I am ", age, " years old")
	fmt.Println("Today is: ", currentYear)

	//Declaring multiple variables
	var (
		color1 = "red"
		color2 = "blue"
		color3 = "green"
	)
	fmt.Println(color1, color2, color3)

	var message, email string
	message = "Hello"
	email = "email@mail.com"
	fmt.Println(message, email)

	var firstName, streetAddress string
	var birthYear

	firstName = "Giovanni"
	streetAddress = "123 Main Street"
	birthYear = 1985
	fmt.Println(firstName, streetAddress, birthYear)
}
