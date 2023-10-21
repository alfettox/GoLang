/*
Author: Giovanni De Franceschi
*/


package main

import "fmt"

func main() {
    var n, m int = 2, 10

    a := n < m    // a is true because 2 is less than 10
    b := n > m    // b is false because 2 is not greater than 10
    c := (n > m) || true // c is true because (2 > 10) is false, but (false || true) is true
    d := (n < m) && true // d is true because (2 < 10) is true, and (true && true) is true

    fmt.Println("a = ", a) // Prints "a = true"
    fmt.Println("b = ", b) // Prints "b = false"
    fmt.Println("c = ", c) // Prints "c = true"
    fmt.Println("d = ", d) // Prints "d = true"
}
