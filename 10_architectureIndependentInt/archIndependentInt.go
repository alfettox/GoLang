/*
	Architecture independent integer types
	- int, uint, uintptr
	- 32 bit systems: 32 bit integers
	- 64 bit systems: 64 bit integers
	- Size of int is platform dependent
	- Size of uint is platform dependent
	- Size of uintptr is platform dependent
	- Use int for indexing a collection
	- Use uint for counting
	- Use uintptr for storing a pointer
	- Use int for integer arithmetic
	- Use uint for bit manipulation
	- Use uintptr to convert between pointers
	- Use int if you need a variable that is guaranteed to be able to hold all the values of a pointer

*/

package main

import "fmt"

func main(){
	var age int8 = 22;
	var port int16 = 2000;
	var zipcode int32 = 500081;
	var phoneNumber int64 = 1234567890; //signed integer
	var phoneNumberUnsigned uint64= 1234567890; //unsigned integer
	var population int = 1000000000;

	fmt.Println(age);
	fmt.Println(port);
	fmt.Println(zipcode);
	fmt.Println(phoneNumber);
	fmt.Println(phoneNumberUnsigned);
	fmt.Println(population);
	

}