/*
Author: Giovanni De Franceschi
*/
package main

import "fmt"

func main() {
	var color string = "resdad"

	switch color {
	case "blue":
		fmt.Println("The color is blue")
	case "red":
		fmt.Println("The color is red")
	case "yellow":
		fmt.Println("The color is yellow")
	default:
		fmt.Println("Invalid color")
	}


}
