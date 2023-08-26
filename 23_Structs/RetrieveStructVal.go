/*
Author: Giovanni De Franceschi
*/


package main

import "fmt"

type account struct {
	number string
	balance float64
 }

func main(){
	var a = account{"C14235345354",140}

	fmt.Println(a)
 
	fmt.Println("The account number is", a.number)
 
	fmt.Print("The current balance is", a.balance)
}