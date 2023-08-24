/*
Author: Giovanni De Franceschi
*/



package main

import (
   "fmt"
)

func main() {
   specialFunc := func(a, b, c int) (x, y, z int){
      x = a + b
      y = a + c
      z = b + c
      return
   }

   fmt.Println(specialFunc(1, 2, 3))
}
