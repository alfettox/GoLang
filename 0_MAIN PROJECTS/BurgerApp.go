/*
Author: Giovanni De Franceschi
*/

package main

import ("fmt")


type Burger struct {
	Bun bool
	Price float64
	Dressed bool
}

type Drink struct {
	Price float64
	Type string
}

type Side struct {
	Price float64
	Type string
}

type Combo struct {
	Burger
	Side
	Drink
	Price float64
}

func orderBurger() Burger {
	fmt.Println("Would you like a burger? (y/n)")
	var burger string
	fmt.Scanln(&burger)
	if burger == "y" {
		return Burger{true, 5.99, true}
	} else {
		return Burger{false, 0, false}
	}
}

func orderDrink() Drink {
	fmt.Println("Would you like a drink? (y/n)")
	var drink string
	fmt.Scanln(&drink)
	if drink == "y" {
		return Drink{1.99, "Coke"}
	} else {
		return Drink{0, ""}
	}
}

func orderSide() Side {
	fmt.Println("Would you like a side? (y/n)")
	var side string
	fmt.Scanln(&side)
	if side == "y" {
		return Side{2.99, "Fries"}
	} else {
		return Side{0, ""}
	}
}

func orderCombo() Combo {
	burger := orderBurger()
	drink := orderDrink()
	side := orderSide()
	return Combo{burger, side, drink, burger.Price + drink.Price + side.Price}
}

func takeOrder() Combo {
	fmt.Println("Welcome to our burger shop!")

	combo := orderCombo()

	fmt.Println("\nThank you for your order!")
	return combo
}


func main(){

	order := takeOrder()

	fmt.Println("\n --- ORDER DETAILS ===")
	fmt.Printf("Burger: %t, Bun: %t, Price: $%.2f\n", order.Burger.Bun, order.Burger.Dressed, order.Burger.Price)
	fmt.Printf("Drink: %s, Price: $%.2f\n", order.Drink.Type, order.Drink.Price)
	fmt.Printf("Side: %s, Price: $%.2f\n", order.Side.Type, order.Side.Price)
	fmt.Println("=====================================")
	fmt.Printf("Combo Price: $%.2f\n", order.Price)


}