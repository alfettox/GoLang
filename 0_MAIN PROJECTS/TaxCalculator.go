/*
Author: Giovanni De Franceschi
*/
package main

import (
	"fmt"
)

func calculateTax(price float64, taxRate float64) float64 {
	return price * taxRate
}

func main() {
	provinceTaxRates := map[string]float64{
		"Ontario":           0.13,
		"Colombie-Britannique": 0.07,
		"Québec":             0.14975,
	}

	selectedProvince := "Ontario" // You can change this to your preferred default province

	fmt.Printf("Selected province: %s\n", selectedProvince)

	for {
		fmt.Println("Select a province:")
		i := 1
		for province := range provinceTaxRates {
			fmt.Printf("%d: %s\n", i, province)
			i++
		}

		var provinceChoice int
		fmt.Print("Enter the number for the province: ")
		fmt.Scanln(&provinceChoice)

		var provinceIndex int
		i = 1
		for province := range provinceTaxRates {
			if i == provinceChoice {
				selectedProvince = province
				provinceIndex = i
				break
			}
			i++
		}

		if provinceIndex == 0 {
			fmt.Println("Invalid province choice.")
			return
		}

		fmt.Print("Enter the price: ")
		var price float64
		fmt.Scanln(&price)

		taxRate := provinceTaxRates[selectedProvince]
		tax := calculateTax(price, taxRate)
		totalPrice := price + tax

		fmt.Printf("Province: %s\n", selectedProvince)
		fmt.Printf("Tax amount: $%.2f\n", tax)
		fmt.Printf("Total price including tax: $%.2f\n", totalPrice)

		var anotherSimulation string
		fmt.Print("Do you want to make another simulation? (yes/no): ")
		fmt.Scanln(&anotherSimulation)
		if anotherSimulation != "yes" {
			break
		}
	}
}
