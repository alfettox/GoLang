/*
Author: Giovanni De Franceschi
*/

package main

import "fmt"

type SavingAccount struct {
   number  string
   balance float64
}

type CheckingAccount struct {
   number  string
   balance float64
}

func (a *SavingAccount) withdraw(value float64) bool {
   if a.balance >= value {
      a.balance = a.balance - value // Corrected subtraction operator
      return true
   }
   return false
}

func (a *CheckingAccount) withdraw(value float64) bool {
   if a.balance >= value {
      a.balance = a.balance - value // Corrected subtraction operator
      return true
   }
   return false
}

func main() {
   a := SavingAccount{}
   a.number = "S21345345345355"
   a.balance = 159

   fmt.Println("withdraw from savings:", a.withdraw(150))
   fmt.Println("new savings balance", a)

   b := CheckingAccount{}
   b.number = "C218678678345345355"
   b.balance = 2000

   fmt.Println("withdraw from checking:", b.withdraw(150))
   fmt.Println("new checking balance", b)
}
