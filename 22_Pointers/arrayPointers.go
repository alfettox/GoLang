/*
Author: Giovanni De Franceschi
*/
package main

import "fmt"

func main() {
    nums := [5]int{1, 2, 3, 4, 5}
    for i := 0; i < len(nums); i++ { // Corrected variable declaration
        fmt.Print(nums[i], " - ")
    }

    var numsPtr [5]*int // Changed the length to match the nums array
	fmt.Println(" ")
    for i := 0; i < len(numsPtr); i++ {
        numsPtr[i] = &nums[i]
    }

    fmt.Println(numsPtr)
}
