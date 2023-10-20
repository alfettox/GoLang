/*
Author: Giovanni De Franceschi
*/






package main

import "fmt"




type account struct {
	number string
	balance float64
	owner entity
 }
 
 type entity struct{
	id string
	address string
 }
 
 func main() {
	e := entity{"000-00-0000", "123 Main Street"}
	a := account{}
	a.number ="C21345345345355"
	a.balance = 140609.09
 
	a.owner = e
 
	fmt.Println(a)
 }