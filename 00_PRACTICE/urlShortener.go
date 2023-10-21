/*
Author: Giovanni De Franceschi
*/

package main

import (
    "fmt"
    "math/rand"
    "strings"
)

func main() {
    var url = "http://www.giovannidefranceschi.com/main/index.html"

    var pageName = url[strings.LastIndex(url, "/")+1:]
    var extension = pageName[strings.LastIndex(pageName, ".")+1:]

    randomDigits := rand.Intn(90) + 10

    shortenedUrl := "http://go.to/" + pageName[:strings.LastIndex(pageName, ".")] + fmt.Sprintf("%02d", randomDigits) + "." + extension

    fmt.Println("Shortened URL:", shortenedUrl)

    myString := "This is my string"

    index := strings.LastIndex(myString, "s")
    finalPart := myString[index:]
    
    fmt.Println(index)
    fmt.Println(finalPart)
}
