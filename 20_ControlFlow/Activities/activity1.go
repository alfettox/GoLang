/*
Author: Giovanni De Franceschi
*/
package main

import (
	"fmt"
)

func main() {
	var money int
	fmt.Print("How much money (in cents) do you have in your pocket? ")
	fmt.Scan(&money)

	if money < 0 {
		fmt.Println("That's not possible!")
	} else if money > 0 && money < 1000 { 
		fmt.Println("You are broke!")
	} else if money > 10000 { 
		fmt.Println("You are rich! Let's go out for a beer!")
	} else {
		fmt.Println("Invalid value provided!")
	}
}
