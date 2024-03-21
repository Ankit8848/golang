package main

import "fmt"

/*  Types

- case-insensitive.
- variable type specify before executing.
- Everything is a Type.

String
Bool
Integer 	uint8	uint64	int8	int64	uintptr 	alias
Floating	float32		float64
Complex

Others:  Arrays	 Slices		Maps	Structs		Pointers	Functions	Channels ...

*/

// for printing  a variable and string
// 1st way to initialize variable

const LoginToken string = "LoveYouAll"

// note: 1st character of variable is capital because we want to tell that it is public.
// so, it can be accessible anywhere.
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
	// var name1 string
	// var isKanjoos bool
	// fmt.Print("Enter your name & are you kanjoos: ")
	// fmt.Scanf("%s %t", &name1, &isKanjoos)
	// fmt.Println(name, isKanjoos)

	//var akj string
	//var bkj int
	//fmt.Print("Enter a string and a number: ")
	//count, err := fmt.Scanf("%s %d", &akj, &bkj)

	// 2nd way to initialize var
	// implicit type

	var friend = "Tabrej Sheikh"
	fmt.Println("\n", friend)

	// 3rd way to initialize var
	// no var style.
	// - inside the method it is allowed to use, but in global and all - not allowed.

	amount := 100
	fmt.Println(amount)

	// To know type:
	fmt.Println(LoginToken)
	fmt.Printf("Variable of type: %T \n", LoginToken)

}
