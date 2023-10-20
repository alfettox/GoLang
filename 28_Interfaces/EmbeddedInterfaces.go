/*
Author: Giovanni De Franceschi
*/




package main

import "fmt"



type AccountOperation interface {
	computeInterest() float64
	displayInformations()
}

type UserOperation interface {
	changeNumber(newNumber string)
}

type BankingOperations struct {
	AccountOperation
	UserOperation
}

type SavingsAccount struct {
	number   string
	balance  float64
	interest float64
}

func (a *SavingsAccount) computeInterest() float64 {
	return 0.001
}

func (a *SavingsAccount) changeNumber(number string) {
	a.number = number
}

func (a *SavingsAccount) displayInformations() {
	fmt.Println("Number:", a.number)
	fmt.Println("Balance:", a.balance)
	fmt.Println("Interest:", a.interest)
}

func main() {
	a := SavingsAccount{}
	a.number = "S21345345345355"
	a.balance = 159

	var ao1 BankingOperations
	ao1.AccountOperation = &a
	ao1.UserOperation = &a

	ao1.displayInformations()
	fmt.Println("ao1 type:")
}
