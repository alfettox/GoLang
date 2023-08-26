/*
Author: Giovanni De Franceschi
*/






package main

import (

	"fmt"
	"strconv"
)

func main(){
	var a int = 20;
	var b int = 10;
	fmt.Println("a = " + strconv.Itoa(a));
	fmt.Println("b = " + strconv.Itoa(b));

	if( a > b){
		goto PART1
	} else {
		goto PART2
	}

	PART1:
		fmt.Println("a is greater than b")
		goto END


	PART2:
		fmt.Println("b is greater than a")
		goto END

	END:
		fmt.Println("-- End of program -- ")

}