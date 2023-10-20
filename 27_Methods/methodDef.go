/*
Author: Giovanni De Franceschi
*/





package main

import "fmt"

// we use the type keyword to create a new type s based on the string type.
type s string

// this method has a receiver of type s as defined above
func(text s) IsEmpty() bool{
   if len(text) > 0{
      return false
   }
   return true
}



func main() {
   text := s("Hi")

   fmt.Println("type value:", text)
   fmt.Println("method value:", text.IsEmpty())
}