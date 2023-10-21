/*
Author: Giovanni De Franceschi
*/
package main

import "fmt"

func main() {
    anonymousFunc := func(a, b, c int) (x, y, z int) {
        x = a + b
        y = a + c
        z = b + c
        return
    }

    fmt.Println(anonymousFunc(1, 2, 3))

    geometricCalc := func(l1, l2, h int) (area, volume int) {
        area = l1 * l2
        volume = area * h
        return
    }

    area, volume := geometricCalc(2, 3, 4)
    fmt.Println("Volume:", volume, "Area:", area)
}
