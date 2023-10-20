/*
Author: Giovanni De Franceschi
*/
package main

import (
    "fmt"
    "strconv"
    "strings"
)

func main() {
    var input string
    fmt.Print("Enter a formatted integer (e.g., 123-456-7890): ")
    fmt.Scan(&input)

    parts := strings.Split(input, "-")

    if len(parts) != 3 {
        fmt.Println("Invalid input format. Please use the format: XXX-XXX-XXXX")
        return
    }

    var part1, part2, part3 int
    var err error

    part1, err = strconv.Atoi(parts[0])
    if err != nil {
        fmt.Println("Invalid input for the first part")
        return
    }

    part2, err = strconv.Atoi(parts[1])
    if err != nil {
        fmt.Println("Invalid input for the second part")
        return
    }

    part3, err = strconv.Atoi(parts[2])
    if err != nil {
        fmt.Println("Invalid input for the third part")
        return
    }

    if len(parts[0]) != 3 || len(parts[1]) != 3 || len(parts[2]) != 4 {
        fmt.Println("Invalid input format. Please use the format: XXX-XXX-XXXX")
        return
    }

    fmt.Printf("Valid number: %03d-%03d-%04d\n", part1, part2, part3)
}
