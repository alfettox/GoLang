/*
Author: Giovanni De Franceschi
*/
package main

import "fmt"


// range return value and index

func main(){
	msg := "Marameo ciao"
	for i, c := range msg{
		fmt.Print(i," ")
		fmt.Println(string(c))
	}
}