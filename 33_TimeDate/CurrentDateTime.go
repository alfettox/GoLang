/*
Author: Giovanni De Franceschi
*/




package main

import (
	"fmt"
	"time"
)


func main (){
	now := time.Now()
	fmt.Println("Current Time:", now)
	fmt.Println("Current Year:", now.Year())
	fmt.Println("Current Month:", now.Month())
	fmt.Println("Current Day:", now.Day())
	fmt.Println("Current Hour:", now.Hour())
	fmt.Println("Current Minute:", now.Minute())
	fmt.Println("Current Second:", now.Second())
	fmt.Println("Current Nanosecond:", now.Nanosecond())
	fmt.Println("Current Weekday:", now.Weekday())
	fmt.Println("Current Year Day:", now.YearDay())  // 240
	fmt.Println("Current Location:", now.Location()) // Local
	fmt.Println("Current weekday: ", now.Weekday())  // Monday
}