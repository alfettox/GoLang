/*
Author: Giovanni De Franceschi
*/




package main

import (
    "fmt"
    "time"
)

func main() {
    for i := 1; i <= 100; i++ {
        duration := time.Duration(i*100) * time.Millisecond
        fmt.Printf("Iteration %d: Waiting for %v\n", i, duration)
        time.Sleep(duration)
    }
}
