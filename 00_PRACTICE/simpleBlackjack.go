/*
Author: Giovanni De Franceschi
*/
package main

import (
    "fmt"
    "math/rand"
)

func main() {
	playerCards := []int{getCard(), getCard()}
	dealerCards := []int{getCard(), getCard()}

    displayHand("Player's Hand:", playerCards)
    fmt.Println("Dealer's Hand:", dealerCards[1])

    for calculateHandValue(playerCards) < 21 {
        fmt.Print("Take or stop? ")
        var choice string
        fmt.Scanln(&choice)
        if choice == "take" {
            playerCards = append(playerCards, getCard())
            displayHand("Player's Hand:", playerCards)
        } else if choice == "stop" {
            break
        }
    }

    for calculateHandValue(dealerCards) < 17 {
        dealerCards = append(dealerCards, getCard())
    }

    playerValue := calculateHandValue(playerCards)
    dealerValue := calculateHandValue(dealerCards)

    fmt.Println("\nPlayer's Hand:")
    displayHand("", playerCards)
    fmt.Printf("Player's Total: %d\n", playerValue)

    fmt.Println("\nDealer's Hand:")
    displayHand("", dealerCards)
    fmt.Printf("Dealer's Total: %d\n", dealerValue)

    switch {
    case playerValue > 21:
        fmt.Println("Dealer wins! Over 21")
    case dealerValue > 21:
        fmt.Println("Player wins!")
    case playerValue == dealerValue:
        fmt.Println("TIE!")
    case playerValue > dealerValue:
        fmt.Println("Player wins!")
    default:
        fmt.Println("Dealer wins!")
    }
}

func getCard() int {
    return rand.Intn(10) + 1
}

func calculateHandValue(hand []int) int {
    total := 0
    for _, card := range hand {
        total += card
    }
    return total
}

func displayHand(title string, hand []int) {
    fmt.Printf("%s %v\n", title, hand)
}
