/*
Author: Giovanni De Franceschi
*/


package main

import (
    "fmt"
)

func main() {
    var value interface{} 

    value = 12          
    fmt.Printf("Value (integer): %v\n", value)   // %v represent the value of a variable

    value = "Hello"   
    fmt.Printf("Value (string): %v\n", value)

	var name string = "Giovanni"
	fmt.Printf("My name is %v", name)
}
