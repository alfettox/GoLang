/*
Author: Giovanni De Franceschi
*/













package main

import(
	"fmt"
)

func main(){
	var a int = 1
	var b int32 = 20
	var c int64 = 30
	var d float32 = 40.5

	// error you need explicit type conversions
	// fmt.Println(a+b)
	// fmt.Println(b+c)
	// fmt.Println(c+d)

	fmt.Println(a+int(b))  //convert b to int
	fmt.Println(b+int32(c)) //convert c to int32
	fmt.Println(c+int64(d)) //convert d to int64


}