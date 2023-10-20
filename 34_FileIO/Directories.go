/*
Author: Giovanni De Franceschi
*/


package main

import (
	"fmt"
	"os"
)

func main() {
	createDir()
	createFile()
	listDir()
}

func createDir() {
	err := os.Mkdir("./newDir", 0777)
	if err != nil {
		panic(err)
	}
}

func createFile() {
	data := "Hello world"
	fmt.Println("Data: ", data)

	f, err := os.Create("./newDir/myNewFile.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	n, err := f.WriteString(data)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bytes written: ", n)
}

func listDir() {
	dir, err := os.Open("./")
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		panic(err)
	}

	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}
