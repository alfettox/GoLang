/*
Author: Giovanni De Franceschi
*/

















package main


import(
	"fmt"

)

var globalVariable string = "This is a GlobalVar"

func main(){
	var email string = "a@b.com"
	fmt.Println(email)
	fmt.Println(globalVariable)
}