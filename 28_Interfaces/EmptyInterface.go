/*
Author: Giovanni De Franceschi
*/





package main

import "fmt"

func main() {
   var s interface{}
   fmt.Println(s)
   fmt.Printf("s is nil and has type %T value %v\n", s, s)
}

//nil is a predeclared identifier representing the zero value for a pointer, channel, func, interface, map, or slice type.