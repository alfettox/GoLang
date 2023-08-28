/*
Author: Giovanni De Franceschi
*/

package main

import "fmt"

type AccountOperations interface{
   // Methods
   computeInterest() float64
}

type SavingsAccount struct {
   number string
   balance float64
   interest float64
}

type CheckingAccount struct {
   number string
   balance float64
   interest float64
}

func(a *SavingsAccount) computeInterest() float64{
   return 0.001
}

func(a *CheckingAccount) computeInterest() float64{
   return 0.005
}

func describe(ao  AccountOperations) {
   // we use %T to display the dynamic type of ao
   // and %v to display the dynamic value
   fmt.Printf("Interface type %T value %v\n", ao, ao)
}

func main() {
   a := SavingsAccount{}
   a.number = "S21345345345355"
   a.balance = 159

   var ao1 AccountOperations
   ao1 = &a
   fmt.Println("ao1 type:")
   describe(ao1)

   b := CheckingAccount{}
   b.number = "C218678678345345355"
   b.balance = 2000

   var ao2 AccountOperations
   ao2 = &b
   fmt.Println("ao2 type:")
   describe(ao2)
}