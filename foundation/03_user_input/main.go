package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	friend := "aman"
	fmt.Println(friend)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your friend height:")

	// 1st way : comma ok || comma err syntax

	input, _ := reader.ReadString('\n')

	// if you dont care about an error, so in place of err - just use "_" .
	// if you dont care about an input, so in place of input - just use "_" .

	fmt.Println("Thanks for saying, ", input)
	fmt.Printf("Type of this friend height is %T", input)

	// 2nd way

	var name string = "ankit"
	var user string = "jha"

	// string formatting => Printf
	// %v => formats the value in default format.
	//fmt.Printf("Template string %s", Object args(s))
	fmt.Printf("Awesome meeting with you %v %v ", name, user)

	// %d => formats decimal integers
	var marks int = 82
	fmt.Printf("Marks: %d", marks)

	// user input
	// Scanf
	// fmt.Scanf("%<format specifier > (s)", object_arguments)
	// var name1 string
	// fmt.Println("Enter Your Name: ")
	// fmt.Scanf("%s", &name1)
	// fmt.Println("Hey there, ", name1)

	// multiple user input
	// var name1 string
	// var isKanjoos bool
	// fmt.Print("Enter your name & are you kanjoos: ")
	// fmt.Scanf("%s %t", &name1, &isKanjoos)
	// fmt.Println(name, isKanjoos)

	//var akj string
	//var bkj int
	//fmt.Print("Enter a string and a number: ")
	//count, err := fmt.Scanf("%s %d", &akj, &bkj)

}
