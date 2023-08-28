/*
Author: Giovanni De Franceschi
*/
package main

import (
	"fmt"
	"os"
)

func main(){
	data := "Hello world"
	fmt.Println("Data: ", data)

	f, error := os.Create("./newFile/myNewFile.txt")
	if(error != nil){
		panic(error)
	}
	defer f.Close()

	n,err := f.WriteString(data)
	if(err != nil){
		panic(err)
	}
	fmt.Println("Bytes written: ", n)
	f.Close()
}