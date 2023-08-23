/*
Author: Giovanni De Franceschi
*/

package main


import "fmt"

func main(){
	var message = "This is a message"
	fmt.Println(message)

	for i, c:= range message { //range returns two values: the index and the character at that index
		fmt.Println("Index =", i, "message[i] =", string(c))
	}
}