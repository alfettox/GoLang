/*
Author: Giovanni De Franceschi
*/
package main

//armstrong number is equal to the sum of its own digits

import "fmt"

func isArmstrongNumber(number int) bool {
    originalNumber := number
    sum := 0

    for number > 0 {
        digit := number % 10
        sum += digit * digit * digit
        number /= 10
    }

    return sum == originalNumber
}

func main() {
    fmt.Println("Armstrong numbers between 1 and 500:")

    for i := 1; i <= 500; i++ {
        if isArmstrongNumber(i) {
            fmt.Println(i)
        }
    }
}
