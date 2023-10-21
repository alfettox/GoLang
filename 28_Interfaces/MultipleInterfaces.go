/*
Author: Giovanni De Franceschi
*/






package main

import (
	"fmt"
)



type AccountOperations interface {
	computeInterest() float64
	displayInformations()
}

type UserOperations interface {
	changeNumber(newNumber string)
}

type SavingAccount struct {
	number   string
	balance  float64
	interest float64
}

func (a *SavingAccount) computeInterest() float64 {
	return 0.001
}

func (a *SavingAccount) changeNumber(number string) {
	a.number = number
}

func (a *SavingAccount) displayInformations() {
	fmt.Println("Number:", a.number)
	fmt.Println("Balance:", a.balance)
	fmt.Println("Interest:", a.interest)
}

func main() {
	a := SavingAccount{}
	a.number = "S21345345345355"
	a.balance = 159

	var ao1 AccountOperations
	var uo1 UserOperations

	ao1 = &a
	uo1 = &a

	uo1.changeNumber("S1234567890")

	fmt.Println("ao1 type:")
	ao1.displayInformations()
}
