/*
Author: Giovanni De Franceschi
*/



package main

import (
   "fmt"
   "reflect"
)

type account struct {
   number string
   balance float64
   owner entity
}

type entity struct{
   id string
   address string
}

func(a account) HaveEnoughBalance(value float64) bool{
   if a.balance >= value{
      return true
   }
   return false
}

func main() {
   e := entity{"555-22-3333","123 Rue Dela"}
   a := account{}
   a.number = "C465466"
   a.balance = 14060.09
   a.owner = e

   fmt.Println("Type and value of a:")
   fmt.Println(reflect.TypeOf(a))
   fmt.Println(reflect.ValueOf(a))

   fmt.Println("\nType and value of e:")
   fmt.Println(reflect.TypeOf(e))
   fmt.Println(reflect.ValueOf(e))
}