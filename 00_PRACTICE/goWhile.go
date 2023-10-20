/*
Author: Giovanni De Franceschi
*/
package main

import (
	"fmt"
	"os"

)

func main() {
	var a int = 0

	for a < 10 {
		if a%2 == 0 {
			fmt.Println(a, "is Even")
		} else {
			fmt.Println(a, "is Odd")
		}
	a++
	}

	os.Exit(0)
}
