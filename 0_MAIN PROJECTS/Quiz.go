/*
Author: Giovanni De Franceschi
*/


package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	questions := []string{
		"What is the capital of Italy?",
		"What is the capital of France?",
		"What is the capital of Germany?",
		"What is the capital of Spain?",
		"What is the capital of Portugal?",
		"What is the capital of Poland?",
	}

	answers := []string{
		"rome",
		"paris",
		"berlin",
		"madrid",
		"lisbon",
		"warsaw",
	}

	var correctAnswers int
	var incorrectAnswers int
	var guess string
	totalQuestions := len(questions)

	for true {
		fmt.Println("Question #", correctAnswers+incorrectAnswers+1, "of", totalQuestions)
		fmt.Println(questions[correctAnswers+incorrectAnswers])
		fmt.Scan(&guess)

		if strings.ToLower(guess) == answers[correctAnswers+incorrectAnswers] {
			fmt.Println("Correct!")
			correctAnswers++
		} else {
			fmt.Println("Incorrect!")
			incorrectAnswers++
		}

		fmt.Println("You have", correctAnswers, "correct answers and", incorrectAnswers, "incorrect answers.")
		fmt.Println("You have", totalQuestions-(correctAnswers+incorrectAnswers), "questions left.")

		if correctAnswers > (totalQuestions / 2) {
			fmt.Println("You win!")
			break
		} else if incorrectAnswers > (totalQuestions / 2) {
			fmt.Println("You lose!")
			break
		}
	}
	fmt.Println("Accuracy: ", float32(correctAnswers)/float32(correctAnswers+incorrectAnswers)*100, "%")
}
