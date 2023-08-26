/*
Author: Giovanni De Franceschi
*/


package main

import "fmt"

func main() {
    var a int = 20
    b := &a

	fmt.Println("a is:", a)

    fmt.Print("Memory address of a: ", b, "\n")

    *b = 99

    fmt.Println("a is now:", a)
}
