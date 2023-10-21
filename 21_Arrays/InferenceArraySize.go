/*
Author: Giovanni De Franceschi
*/








package main

import "fmt"

func main() {
    numbers := [...]int{1, 2, 4, 5, 6, 7, 8, 9}

    fmt.Println("numbers:", numbers)
    fmt.Println("Length of the numbers array is", len(numbers))

    newArray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    fmt.Println("newArray:", newArray)

    newArray2 := [4]int{} // Initialize with default zero values
    fmt.Println("newArray2:", newArray2)
}
