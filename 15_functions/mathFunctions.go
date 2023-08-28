/*
Author: Giovanni De Franceschi
*/













package main

import (
	"fmt"
	"math"
)

func main(){
	var a,b,c float64 = 3.2,4.5,5.6
    fmt.Println("a =", a, "\nb =", b, "\nc =", c)

	fmt.Println("math.Abs(a) =", math.Abs(a))     // absolute value
	fmt.Println("math.Ceil(b) =", math.Ceil(b))   // round up
	fmt.Println("math.Floor(c) =", math.Floor(c)) // round down
	fmt.Println("math.Max(a,b) =", math.Max(a,b)) // maximum
	fmt.Println("math.Min(a,b) =", math.Min(a,b)) // minimum
	fmt.Println("math.Pow(a,b) =", math.Pow(a,b)) // power
	fmt.Println("math.Sqrt(b) =", math.Sqrt(b))   // square root


}