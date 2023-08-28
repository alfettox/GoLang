/*
Author: Giovanni De Franceschi
*/
















package main

import (
	"fmt"
	"reflect"
)

func main(){
	email := "email@mail.com"
	fmt.Println(email)

	fmt.Println("Type of email variable is: ", reflect.TypeOf(email)) //show the type

	valueNumber := 123
	fmt.Println("Type of valueNumber: ", reflect.TypeOf(valueNumber))
}