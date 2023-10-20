/*
Author: Giovanni De Franceschi
*/






package main

import "fmt"

type account struct {
   number string
   balance float64
}

// This a method defined on the struct account. The receiver argument
// is (a account) which is separate from the input argument list (value).
func(a account) HaveEnoughBalance(value float64) bool {
   if a.balance >= value{
      return true
   }
   return false
}

// this is a simple function
func HaveEnoughBalance2(a account, value float64) bool {
   if a.balance >= value{
      return true
   }
   return false
}

func main() {
   a := account{}
   a.number = "C21345345345355"
   a.balance = 0

   // call the method defined on account
   fmt.Println("Method result:", a.HaveEnoughBalance(150))

   // call the function
   fmt.Println("Function result:", HaveEnoughBalance2(a, 150)) 
}