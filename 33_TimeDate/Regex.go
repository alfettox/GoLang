/*
Author: Giovanni De Franceschi
*/
package main

import (
   "fmt"
   "regexp"
)

func main() {
   // check if string starts with C and ends with n
   m1,err1 := regexp.MatchString("C([a-z]+)n", "Catelyn") 
   fmt.Println(m1)
   fmt.Println("ERROR:",err1)

   // check if string contains at least one digit
   m2,err2 := regexp.MatchString("[0-9]", "jonathan6smith") 
   fmt.Println(m2)
   fmt.Println("ERROR:",err2)
}