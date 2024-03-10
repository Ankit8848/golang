package main

import "fmt"

// printing  a variable and string
func main() {
	var name string = "ankit"
	var user string = "jha"
	fmt.Println("welcome here ", name, ", ", user)
	// new line character => \n
	fmt.Print(name, "\n")
	fmt.Print(user, "\n")

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
	var name1 string
	var isKanjoos bool
	fmt.Print("Enter your name & are you kanjoos: ")
	fmt.Scanf("%s %t", &name1, &isKanjoos)
	fmt.Println(name, isKanjoos)

	//var akj string
	//var bkj int
	//fmt.Print("Enter a string and a number: ")
	//count, err := fmt.Scanf("%s %d", &akj, &bkj)
}
