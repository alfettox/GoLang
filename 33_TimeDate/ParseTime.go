/*
Author: Giovanni De Franceschi
*/
package main

import (
	"fmt"
	"time"
)

func main(){
	t1, e := time.Parse(time.RFC3339, "2018-08-31T10:00:00+07:00")
	if e != nil {
		fmt.Println(e)
	}

	fmt.Println(t1)
	fmt.Println(t1.Day())
	fmt.Println(t1.Month())

	
}