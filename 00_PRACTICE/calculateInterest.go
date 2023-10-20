/*
Author: Giovanni De Franceschi
*/
package main

import "fmt"
import "math"

func main() {
    var P float64
    var r float64
    var n int
    var t int

    fmt.Print("Enter the principal amount (P): ")
    fmt.Scan(&P)

    fmt.Print("Enter the annual interest rate (r): ")
    fmt.Scan(&r)

    fmt.Print("Enter the number of times interest is compounded per year (n): ")
    fmt.Scan(&n)

    fmt.Print("Enter the number of years (t): ")
    fmt.Scan(&t)

    V := P * math.Pow(1+r/float64(n), float64(n*t))

    fmt.Printf("The future value of the investment is: %.2f\n", V)
}
