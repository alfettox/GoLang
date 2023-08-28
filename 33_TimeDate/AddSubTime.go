/*
Author: Giovanni De Franceschi
*/

package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("Current Time:", now)

	// Create a custom time instance corresponding to August 1, 2023, at 12:00:00 UTC.
	customTime := time.Date(2023, time.August, 1, 12, 0, 0, 0, time.UTC)
	fmt.Println("Custom Time:", customTime)

	// Subtract the custom time from the current time.
	diff := now.Sub(customTime)
	fmt.Println("Difference:", diff)
	fmt.Println("Difference in hours:", diff.Hours())
	fmt.Println("Difference in minutes:", diff.Minutes())
	fmt.Println("Difference in seconds:", diff.Seconds())

	// Add the time difference back to the current time.
	add := now.Add(diff)
	fmt.Println("Add:", add)
	fmt.Println("Add in hours:", add.Hour())
	fmt.Println("Add in minutes:", add.Minute())
	fmt.Println("Add in seconds:", add.Second())

	rounded := math.Round(add.Second())
	fmt.Println("Rounded:", rounded)

}
