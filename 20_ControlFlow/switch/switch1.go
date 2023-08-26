/*
Author: Giovanni De Franceschi
*/








package main

import (
	"fmt"
	"strings" // Import the strings package for ToLower()
)

func main() {
	fmt.Print("Enter OS: ")
	var os string
	fmt.Scan(&os)
	os = strings.ToLower(os) // Use strings.ToLower() to convert to lowercase
	switch os {
	case "linux":
		fmt.Println("Nice! You are a Linux user")
	case "windows":
		fmt.Println("Cool! You are a Windows user")
	case "mac os", "mac", "macos": // Separate each value with a comma
		fmt.Println("You are a macOS user")
	default:
		fmt.Println("You are using another OS")
	}
}
