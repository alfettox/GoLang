/*
Author: Giovanni De Franceschi
*/


package main

import "fmt"

type AccountOperations interface{
   // methods
   withdraw(value float64) bool
   deposit(value float64) bool
   displayInfo()
}

type account struct {
   number string
   balance float64
}

func (a *account) withdraw(value float64) bool {
   if a.balance >= value{
      a.balance = a.balance - value
      return true
   }
   return false
}

func (a *account) deposit(value float64) bool {
   if value > 0{
      a.balance = a.balance + value
      return true
   }
   return false
}

func (a *account) displayInfo() {
   fmt.Println(a.balance)
   fmt.Println(a.number)
}

func main() {
   var ao AccountOperations
   fmt.Println("initial value:",ao) 

    // assign ao a pointer to an account value that we created. We can only do this because the account type implements all the methods of AccountOperations Interface.
   ao = &account{"C13443533535",1500}

   //withdrawal amount
   ao.withdraw(150)
   ao.displayInfo()

   // deposit amount
   ao.deposit(1000)
   ao.displayInfo()
}
