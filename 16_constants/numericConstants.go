/*
Author: Giovanni De Franceschi
*/
















package main

import	(
	"fmt"
)

const (
	Big  = 1 << 100 // shift 1 bit left 100 places
	Small = Big >> 99 // shift it right again 99 places, so we end up with 1 << 1, or 2
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat (x float64) float64 {
	return x * 0.1
}

func main(){
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}