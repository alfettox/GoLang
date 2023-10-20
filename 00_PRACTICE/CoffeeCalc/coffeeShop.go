/*
Author: Giovanni De Franceschi
*/
package main

import (
	"fmt"
)

type Coffee struct {  //like a java class
	SizeCost    int
	TypeCost    int
	FlavorCost  int
	TotalCost   int
}

func main() {
	var coffee Coffee

	fmt.Print("Enter the size of the coffee (small, medium, large): ")
	var size string
	fmt.Scanln(&size)
	switch size {
	case "small":
		coffee.SizeCost = 2
	case "medium":
		coffee.SizeCost = 3
	case "large":
		coffee.SizeCost = 4
	default:
		fmt.Println("Invalid size")
		return
	}

	fmt.Print("Enter the type of coffee (regular, decaf, cappuccino): ")
	var coffeeType string
	fmt.Scanln(&coffeeType)
	switch coffeeType {
	case "regular":
		coffee.TypeCost = 1
	case "decaf":
		coffee.TypeCost = 2
	case "cappuccino":
		coffee.TypeCost = 3
	default:
		fmt.Println("Invalid coffee type")
		return
	}

	fmt.Print("Enter the flavor of coffee (none, vanilla, caramel): ")
	var flavor string
	fmt.Scanln(&flavor)
	switch flavor {
	case "none":
		coffee.FlavorCost = 0
	case "vanilla":
		coffee.FlavorCost = 1
	case "caramel":
		coffee.FlavorCost = 2
	default:
		fmt.Println("Invalid flavor")
		return
	}

	coffee.TotalCost = coffee.SizeCost + coffee.TypeCost + coffee.FlavorCost
	fmt.Printf("Total cost of your custom coffee: $%d\n", coffee.TotalCost)
}
