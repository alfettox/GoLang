/*
Author: Giovanni De Franceschi
*/








package main
import "fmt"

type account struct {
	name string
	balance float64
}


func main() {
	var account1 = account{"John", 1000.50}

	fmt.Println(account1)
}