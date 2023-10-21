/*
Author: Giovanni De Franceschi
*/
package main

import ("fmt"
		)


func main(){
	arr := [...] string {"mario", "emma", "chiara"}
	fmt.Println(arr)


	numbers := [...] int {1,3,4,6}

	fmt.Println(numbers)

	for i, value := range numbers{
		fmt.Print("i:",i)
		fmt.Print(" ", value,"\n")
	}

	numbers[2] = 156

	fmt.Println(numbers)
}