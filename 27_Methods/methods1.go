/*
Author: Giovanni De Franceschi
*/

package main

import (
	"fmt"
)

func main() {
	e := entity{"000-0000-0001", "123 rue Sherbrooke"}
	a := account{
		number:  "123456",
		balance: 1000.00,
		owner:   e,
	}

	fmt.Println("Account Information:")
	fmt.Println("Account Number:", a.number)
	fmt.Println("Balance:", a.balance)
	fmt.Println("Owner ID:", a.owner.id)
	fmt.Println("Owner Address:", a.owner.address)

	// Check if the account has sufficient balance
	value := 500.00
	if a.SufficientBalance(value) {
		fmt.Printf("The account has sufficient balance for a withdrawal of %.2f\n", value)
	} else {
		fmt.Printf("The account does not have sufficient balance for a withdrawal of %.2f\n", value)
	}
}

type account struct {
	number  string
	balance float64
	owner   entity
}

type entity struct {
	id      string
	address string
}

func (a account) SufficientBalance(value float64) bool {
	return a.balance >= value
}
