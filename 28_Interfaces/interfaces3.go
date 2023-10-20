/*
Author: Giovanni De Franceschi
*/




package main

import "fmt"



type AccountOperationsInterface interface {
	computeInterest() float64
}

type SavingsAccount struct {
	Number   string
	Balance  float64
	Interest float64
}

type CheckingAccount struct {
	Number   string
	Balance  float64
	Interest float64
}

func (a *SavingsAccount) computeInterest() float64 {
	return a.Balance * a.Interest
}

func (a *CheckingAccount) computeInterest() float64 {
	return a.Balance * a.Interest
}

func main() {
	savings := &SavingsAccount{
		Number:   "S1234",
		Balance:  1000,
		Interest: 0.05,
	}

	var accountOperation1 AccountOperationsInterface
	accountOperation1 = savings
	fmt.Println("Savings interest:", accountOperation1.computeInterest())

	checking := &CheckingAccount{
		Number:   "C1235",
		Balance:  3000,
		Interest: 0.03,
	}

	var accountOperation2 AccountOperationsInterface
	accountOperation2 = checking
	fmt.Println("Checking interest:", accountOperation2.computeInterest())
}
