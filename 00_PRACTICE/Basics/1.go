/*
Author: Giovanni De Franceschi
*/

package main

import (
    "fmt"
    "os"
    "strconv"
)


var globalVar int = 3

func main() {
    fmt.Println("This is a print")

    fmt.Println(globalVar)

    printSome()

    var msg string = "This is a string"

    fmt.Println(msg)
    fmt.Println(&msg)

    msg2 := "This is a dynamically typed string"

    fmt.Println(msg2)
    fmt.Println(&msg2)

    var input1, input2 string
    fmt.Println("Enter two strings:")
    fmt.Scan(&input1)
    fmt.Scan(&input2)

    fmt.Println(input1)
    fmt.Println(&input1)
    fmt.Println(input2)
    fmt.Println(&input2)

    // a, b, c := "anna", "panna", "capanna"

    // fmt.Println(a)
    // fmt.Println(b)
    // fmt.Println(c)

    var message, year = "Hello world", 2022

    fmt.Println(message)
    fmt.Println(year)

    var inputInt1, inputInt2 string
    fmt.Println("Enter two integers:")

    fmt.Scan(&inputInt1)
    fmt.Scan(&inputInt2)

    intInput1, _ := strconv.Atoi(inputInt1)
    intInput2, _ := strconv.Atoi(inputInt2)

    intInputNew, error := strconv.Atoi("mannaggialapeppa")

    fmt.Println(intInputNew, error)

    sum := intInput1 + intInput2

    fmt.Println("Sum:", sum)


    fmt.Println(globalVar)


    var staticTypeVar int = 123


    fmt.Println(staticTypeVar)

    staticTypeVar = 1

    fmt.Println(staticTypeVar)

    // staticTypeVar = 123.2323

    // fmt.Println(staticTypeVar)

    dynamicTypeToBeReassigned := 123
    fmt.Println(dynamicTypeToBeReassigned)
    // dynamicTypeToBeReassigned = 1.2  ERROR GO IS A STATICALLY TYPED LANGUAGE
    // fmt.Println(dynamicTypeToBeReassigned)


    var a, b, c int32 = 20,10,8
    fmt.Print("a = ", a, "\n")
    fmt.Print("b = ", b, "\n")
    fmt.Print("c = ", c, "\n")
    fmt.Print("a + b = ", a+b, "\n")
    fmt.Print("a - b = ", a-b, "\n")
    fmt.Print("a *b = ", a*b, "\n")
    fmt.Print("a / b= ", a/b, "\n")
    fmt.Print("a mod b = ", a%b, "\n")

    // int 2^31  32 bit 
    // int platform specific, can be 32 in 32bit and 64 in 64but systems 

    // var t int8 = 100
    // var y int = 70

    // fmt.Println("t =", t)
    // fmt.Println("y =", y)
    // fmt.Println("t == y:", t == y) // CANT COMPARE INT AND INT8
    // fmt.Println("t != y:", t != y) 

    // fmt.Println("t =", t)

    var t int = 100
    var y int = 70

    fmt.Println("t =", t)
    fmt.Println("y =", y)
    fmt.Println("t == y:", t == y)
    fmt.Println("t != y:", t != y) 


    


    os.Exit(0)
}

func printSome(){
    fmt.Println(globalVar)
    globalVar +=1
}