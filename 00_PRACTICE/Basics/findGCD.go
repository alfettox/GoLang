/*
Author: Giovanni De Franceschi
*/

package main

import "fmt"

func findGCD(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

func main() {
    var num1, num2 int

    fmt.Print("Enter the first number: ")
    fmt.Scan(&num1)

    fmt.Print("Enter the second number: ")
    fmt.Scan(&num2)

    result := findGCD(num1, num2)

    fmt.Println("The GCD is:", result)
}
