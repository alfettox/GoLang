/*
Author: Giovanni De Franceschi
*/
package main


import ("fmt"
		"reflect")




func main() {
	var array = [5] int {1,2,3,4,5}

	fmt.Println("Array: ", array)
	fmt.Println("Array type: ", reflect.TypeOf(array)) //reflect.TypeOf() returns the type of a variable

	y := array[1:4] //slice of the array, as in Python
	fmt.Println("Slice: ", y)
	fmt.Println("Slice type: ", reflect.TypeOf(y))

    numbers := []int{0, 1, 2, 5, 798, 43, 78}

    s := numbers[1:5] //length 4, capacity 6
    fmt.Println("Slice s:", s)
    fmt.Println("Length of slice s:", len(s))
    fmt.Println("Capacity of slice s:", cap(s))

	//capacity is from the start of the slice to the end of the underlying array
}

