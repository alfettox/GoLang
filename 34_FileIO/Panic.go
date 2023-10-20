/*
Author: Giovanni De Franceschi
*/



package main

import (
	"fmt"
	// "io/ioutil"
	"os"
)

func main(){

	_, error := os.Open("non-existing-file.txt")
	if(error != nil){
		panic(error)
	}
	fmt.Println("After panic")
}