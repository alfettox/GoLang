/*
Author: Giovanni De Franceschi
*/


package main

import (
	"fmt"
	"strings"
)

func main(){
	name := "Charlotte"

	displayUpper(name);
}

func displayUpper(name string){
	fmt.Println(strings.ToUpper(name))
}