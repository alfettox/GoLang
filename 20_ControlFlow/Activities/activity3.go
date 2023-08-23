/*
Author: Giovanni De Franceschi
*/



package main

import (
	"fmt"
	"strings"
)

func main() {
	questions := []string{
		"Is the sky blue? (True/False)",
		"Is water wet? (True/False)",
		"Is the Earth flat? (True/False)",
	}

	correctAnswers := []bool{true, true, false}
	userAnswers := make([]bool, len(questions))
	correctCount := 0

	for i, question := range questions {
		fmt.Println(question)
		var userResponse string
		fmt.Scan(&userResponse)

		userResponse = strings.ToLower(strings.TrimSpace(userResponse))
		if userResponse == "true" || userResponse == "t" {
			userAnswers[i] = true
		} else if userResponse == "false" || userResponse == "f" {
			userAnswers[i] = false
		} else {
			fmt.Println("Invalid response. Please enter True or False.")
			i--
			continue
		}

		if userAnswers[i] == correctAnswers[i] {
			correctCount++
		}
	}

	fmt.Println("\nResults:")
	for i, question := range questions {
		result := "Incorrect"
		if userAnswers[i] == correctAnswers[i] {
			result = "Correct"
		}
		fmt.Printf("%s\nCorrect Answer: %v | Your Answer: %v\n", question, correctAnswers[i], result)
	}

	accuracy := float64(correctCount) / float64(len(questions)) * 100
	fmt.Printf("\nCorrect Response Rate: %.2f%%\n", accuracy)
}
