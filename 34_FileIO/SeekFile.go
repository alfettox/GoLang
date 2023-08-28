/*
Author: Giovanni De Franceschi
*/
package main

import (
	"fmt"
	"os"
)


func main(){
	file, error := os.Open("./flatland01.txt")
	if(error != nil){
		panic(error) // This will cause the program to exit
	}
	s, err := file.Seek(100, 0)
	fmt.Println(s, err)
	data := make([]byte, 50)
	n, err := file.Read(data)
	
	fmt.Println(err);
	fmt.Println("Bytes read:", n)
	fmt.Println("Reading starting from 100th byte:", string(data[:n]))
	file.Close()



}