/*
Author: Giovanni De Franceschi
*/









package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var grandTotal float64 = 0.00

	coffeeType := []string{"Espresso", "Latte", "Cappuccino", "Mocha", "Americano"}
	coffeeSize := []string{"Small", "Medium", "Large"}
	coffeeFlavor := []string{"None", "Vanilla", "Caramel", "Hazelnut", "Mocha"}

	typePrice := map[string]float64{
		"Espresso":   2.99,
		"Latte":      4.50,
		"Cappuccino": 3.75,
		"Mocha":      4.75,
		"Americano":  2.25,
	}

	sizePrice := map[string]float64{
		"Small":  0.00,
		"Medium": 0.50,
		"Large":  1.00,
	}

	flavorPrice := map[string]float64{
		"None":     0.00,
		"Vanilla":  0.50,
		"Caramel":  0.50,
		"Hazelnut": 0.50,
		"Mocha":    0.75,
	}

	fmt.Println("*************************************")
	fmt.Println("* Welcome to Giovanni's Coffee Shop *")
	fmt.Println("*************************************")

	for {
		var orderType string
		listOptions("What would you like to order?", coffeeType, typePrice)
		orderType = validateChoice(coffeeType)

		var orderSize string
		listOptions("What size would you like?", coffeeSize, sizePrice)
		orderSize = validateChoice(coffeeSize)

		var orderFlavor string
		listOptions("Would you like to add a flavor?", coffeeFlavor, flavorPrice)
		orderFlavor = validateChoice(coffeeFlavor)

		coffeeCost := calculateCost(orderType, orderSize, orderFlavor, typePrice, sizePrice, flavorPrice)
		grandTotal += coffeeCost

		fmt.Println("Your order is:", orderSize, orderType, orderFlavor)
		fmt.Println("Your TOTAL for this order is:", coffeeCost, "$")
		fmt.Println("Your GRAND TOTAL so far is:", grandTotal, "$")

		fmt.Println("Would you like to order again? (yes/no)")
		var continueOrder string
		fmt.Scanln(&continueOrder)
		if continueOrder == "no" {
			break
		}
	}
}

func calculateCost(coffeeType, coffeeSize, coffeeFlavor string, typePrice, sizePrice, flavorPrice map[string]float64) float64 {
	totalCost := typePrice[coffeeType] + sizePrice[coffeeSize] + flavorPrice[coffeeFlavor]
	return totalCost
}

func listOptions(prompt string, options []string, prices map[string]float64) {
	fmt.Println(prompt)
	for _, option := range options {
		fmt.Printf("> %s", option)
		if prices != nil {
			price := prices[option]
			fmt.Printf("%*s$%.2f\n", 25-len(option), "", price)
		} else {
			fmt.Println()
		}
	}
	fmt.Print("\nEnter the name: ")
}

func validateChoice(options []string) string {
	reader := bufio.NewReader(os.Stdin)
	for {
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		for _, option := range options {
			if choice == option {
				return choice
			}
		}
		fmt.Println("Invalid choice. Please retry.")
		fmt.Print("Enter your choice: ")
	}
}
