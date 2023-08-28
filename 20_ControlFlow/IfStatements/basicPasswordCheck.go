/*
Author: Giovanni De Franceschi
*/











package main

import (
	"fmt"
)

func main(){
	var username string
	fmt.Println("Enter your username: ")
	fmt.Scan(&username)
	fmt.Println("Enter a password. Check if it's correct or not")
	var password string
	fmt.Scan(&password)

	passwordChecker(password, username)
}

func passwordChecker(password string, username string){
	var correctPassword string = "password"
	var correctUsername string = "root"
	if (password == correctPassword && username == correctUsername){
		fmt.Println("Username:",username,".Correct password. \nAccess granted")
	} else {
		fmt.Println("Incorrect password")
	}
}