/*
Author: Giovanni De Franceschi
*/



















package main

import "fmt"

func main() {
    var a, b, c int = 10, 20, 30
    fmt.Println("a=", a)
    fmt.Println("b=", b)
    fmt.Println("c=", c)

    var d, e, f, g, h float32

    d = float32(a + b*c) // Convert the result to float32
    fmt.Println("d=", d)
    
    e = float32((a + b) * c) // Convert the result to float32
    fmt.Println("e=", e)
    
    f = float32(a + b%c) // Convert the result to float32
    fmt.Println("f=", f)
    
    g = float32(a * b / c) // Convert the result to float32
    fmt.Println("g=", g)
    
    h = float32(a) * (float32(b) / float32(c)) // Convert variables and result to float32
    fmt.Println("h=", h)
}
