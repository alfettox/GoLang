/*
Author: Giovanni De Franceschi
*/




package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the current time.
	now := time.Now()
	fmt.Println("Current Time:", now)

	// Create a custom time instance corresponding to January 1, 2018, at 12:00:00 UTC.
	customTime := time.Date(2018, time.January, 1, 12, 0, 0, 0, time.UTC)
	fmt.Println("Custom Time:", customTime)

	// Compare the custom time with the current time.
	fmt.Println("The custom time is after now:", customTime.After(now))
	fmt.Println("The custom time is equal to now:", customTime.Equal(now))
}
