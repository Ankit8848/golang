package main

import "fmt"

func main() {
	fmt.Println("Hey there, I am pointers")

	//var ptr *int
	//
	//fmt.Println("Value of pointer is ", ptr)
	// if we don't specify value then it will show nil.

	numb := 13

	var ptr = &numb
	fmt.Println("Address Value of actual pointer is ", ptr)
	fmt.Println("Value of actual pointer is ", *ptr)

	*ptr = *ptr * 5
	fmt.Println("New value is: ", numb)
	// so here actual change happens on original copy

}
