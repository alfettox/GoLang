/*
Author: Giovanni De Franceschi
*/







package main

import "fmt"

func main() {
    var numbers [5]int

    for a := 0; a < len(numbers); a++ {
        numbers[a] = a + 1
    }

    fmt.Println("numbers:", numbers)

    var newArray [10]int

    for a := 0; a < len(newArray); a++ {
        newArray[a] = a + 1
    }

    fmt.Println("newArray:", newArray)
}
