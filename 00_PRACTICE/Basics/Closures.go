/*
Author: Giovanni De Franceschi
*/
package main

import (
	"fmt"
)

func main(){
	caio := 2
	b := 3

	add := func() int {		//lambda
		return caio + b
	}

	fmt.Println(add())
}