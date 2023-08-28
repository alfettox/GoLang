/*
Author: Giovanni De Franceschi
*/


package main

import "fmt"

type Shape interface {	// Shape is an interface
    Area() float64      // Area is a method
}

type Circle struct {    // Circle is a struct 
    Radius float64
}

func (c Circle) Area() float64 {        // Area() is a method of Circle, which implements Shape
    return 3.14 * c.Radius * c.Radius
}

type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func main() {
    var s Shape
    s = Circle{Radius: 5}
    fmt.Println("Circle area:", s.Area())

    s = Rectangle{Width: 3, Height: 4}
    fmt.Println("Rectangle area:", s.Area())
}
