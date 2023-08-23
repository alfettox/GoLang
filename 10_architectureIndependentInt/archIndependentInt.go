/*
Author: Giovanni De Franceschi
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