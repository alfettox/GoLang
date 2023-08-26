/*
Author: Giovanni De Franceschi
*/


package main

import "fmt"

type account struct {
	name    string
	balance float64
}

func main() {
	var account1 = account{"John", 1000.50}
	fmt.Println(account1)

	var account2 = account{name: "Mary", balance: 2000.50}

	fmt.Println(account2)

}