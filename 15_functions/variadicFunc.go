/*
Author: Giovanni De Franceschi
*/



package main


import (
	"fmt"
)


func sumNumbers( numbers ... int){
	sum := 0
	for i, num := range numbers{  //i is the index, num is the value
		fmt.Println("Current element: ", num, " Current index: ", i)
		sum += num

	}

	fmt.Println("Sum of numbers: ", sum)
	return
}

func main (){
	sumNumbers (1,2,3,4,5,6,7,8,9)
	sumNumbers (2,5,7)
}

