/*
Author: Giovanni De Franceschi
*/




package main

import "fmt"

func main() {
    var a int = 10
    b := &a
    c := &a

    fmt.Println("a is:", a)
    fmt.Println("b is:", b)
    fmt.Println("c is:", c)
    var x *int

    fmt.Print("Memory address of a: ", b, "\n")
    fmt.Print("Memory address of a: ", c, "\n")
    fmt.Print("Memory address of a: ", x, "\n")

	fmt.Print("Is c == b? ")
	fmt.Println(c == b) // compare c and b
	
	fmt.Print("Is x == b? ")
	fmt.Println(x == b) // compare x and b
}
