/*
Author: Giovanni De Franceschi
*/

package main

import (
	"fmt"
	"strconv"
	"os"
)

func main() {
	for a := 1; a <= 10; a++ {
		if a%2 == 0 {
			fmt.Println(strconv.Itoa(a) + " is even")
		} else {
			fmt.Println(strconv.Itoa(a) + " is odd")
		}
	}

	var b int = 1

	fmt.Print("\n")

	for ; b <= 10;{


		if(b % 2 ==0){
			fmt.Println(strconv.Itoa(b) + " is a even number")
		}
		b++
	}

		os.Exit(0)

}
