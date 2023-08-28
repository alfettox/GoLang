/*
Author: Giovanni De Franceschi
*/
package main

import (
	"fmt"
	// "io/ioutil"
	"os"
)

func main() {
	// Get the current working directory
	// cwd, err := os.Getwd()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Println("Current working directory:", cwd)

	// Change the current directory to another directory
	// newDir := "./output/"
	// err = os.Chdir(newDir)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Println("Changed to:", newDir)
	filePath := "./flatland01.txt"

	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Print(string(data))
	
	// Write to a file
	outputFilePath := "./output/output.txt" // Use a different variable name
	err = os.WriteFile(outputFilePath, data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}
