/*
Author: Giovanni De Franceschi
*/


package main

import(
	"fmt"
	"os"
	"bufio"
)


func main(){
	file, error := os.Open("./flatland01.txt")
	if(error != nil){
		fmt.Println(error)
		return
	}
	defer file.Close()
	
	scanner:= bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan(){
		lines = append(lines, scanner.Text())
	}

	file.Close()

	for _, line := range lines{
		fmt.Println("line: ",line)
	}

}