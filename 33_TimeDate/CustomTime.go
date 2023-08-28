/*
Author: Giovanni De Franceschi
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a custom time instance representing January 1, 2018, at 12:00:00 UTC.
	customTime := time.Date(2018, time.January, 1, 12, 0, 0, 0, time.UTC)

	// Print various components of the custom time instance.
	fmt.Println("Custom Time:", customTime)
	fmt.Println("Custom Year:", customTime.Year())
	fmt.Println("Custom Month:", customTime.Month())
	fmt.Println("Custom Day:", customTime.Day())
	fmt.Println("Custom Hour:", customTime.Hour())
	fmt.Println("Custom Minute:", customTime.Minute())
	fmt.Println("Custom Second:", customTime.Second())
	fmt.Println("Custom Nanosecond:", customTime.Nanosecond())
	fmt.Println("Custom Weekday:", customTime.Weekday())
	fmt.Println("Custom Year Day:", customTime.YearDay())
	fmt.Println("Custom Location:", customTime.Location())
	fmt.Println("Custom Weekday:", customTime.Weekday())
}
