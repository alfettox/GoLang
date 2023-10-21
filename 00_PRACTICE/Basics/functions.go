/*
Author: Giovanni De Franceschi
*/
package main

import (
    "fmt"
    "strings"
    "bufio"
    "os"
)

func main() {
    var a string
    fmt.Print("Enter a text: ")
    reader := bufio.NewReader(os.Stdin)
    a, _ = reader.ReadString('\n')

    result := DisplayCapitalized(a)
    fmt.Println(result)
}

func DisplayCapitalized(a string) string {
    return strings.ToUpper(a)
}
