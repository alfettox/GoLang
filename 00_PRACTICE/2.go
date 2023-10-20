/*
Author: Giovanni De Franceschi
*/
package main


import("fmt";"math"
		)


func main(){
	var a, b, c float64 =  -23,23.3, 2.8

	fmt.Println(math.Abs(a))
	fmt.Println(math.Ceil(b))
	fmt.Println(math.Floor(c))

	fmt.Println(math.Exp(a))
	fmt.Println(math.Sqrt(c))
	fmt.Println(math.Trunc(b))
	fmt.Println(math.Round(b))
	fmt.Println(math.Pow(a,math.Ceil(b))) 

}