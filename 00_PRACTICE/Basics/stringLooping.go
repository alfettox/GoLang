/*
Author: Giovanni De Franceschi
*/


package main


import "fmt"


func main()	{
	var message string = "Hello world"

	for i := 0 ; i < len(message); i++{
		fmt.Println(string(message[i]))
	}
}