/*
Author: Giovanni De Franceschi
*/

package main

import "fmt"

func main() {
    var a, b int16 = 1, 2
    fmt.Println("a =", a)
    fmt.Println("b =", b)

    // Binary AND
    resultAND := a & b
    fmt.Printf("a & b: %04b & %04b = %04b\n", a, b, resultAND) // a & b: 0001 & 0010 = 0000

    // Binary OR
    resultOR := a | b
    fmt.Printf("a | b: %04b | %04b = %04b\n", a, b, resultOR) // a | b: 0001 | 0010 = 0011

    // Binary XOR
    resultXOR := a ^ b
    fmt.Printf("a ^ b: %04b ^ %04b = %04b\n", a, b, resultXOR) // a ^ b: 0001 ^ 0010 = 0011

    // Binary Left Shift
    resultLeftShift := a << b
    fmt.Printf("a << b: %04b << %04b = %04b\n", a, b, resultLeftShift) // a << b: 0001 << 0010 = 0100

    // Binary Right Shift
    resultRightShift := a >> b
    fmt.Printf("a >> b: %04b >> %04b = %04b\n", a, b, resultRightShift) // a >> b: 0001 >> 0010 = 0000
}
