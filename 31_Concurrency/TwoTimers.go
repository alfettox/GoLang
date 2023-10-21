/*
Author: Giovanni De Franceschi
*/



package main

import (
   "fmt"
   "time"
)

func goroutine(limit int) {
   for i := 0;i < limit; i++ {
      time.Sleep(250 * time.Millisecond)
      fmt.Println("calling goroutine")
      fmt.Println(i)
   }
}

func anothergoroutine(limit int) {
   for i := 0;i < limit; i++ {
      time.Sleep(400 * time.Millisecond)
      fmt.Println("calling anothergoroutine")
      fmt.Println(i)
   }
 }

func main() {
   go goroutine(10)
   go anothergoroutine(10)

   time.Sleep(6 * time.Second)
   fmt.Println("main goroutine")
}