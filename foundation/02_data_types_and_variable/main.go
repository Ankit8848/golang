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

// Go is statically typed.
// variables are assigned a type, either explicitly or implicitly.

// syntax: var <variable name> <data type> = <Value>

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

/*
	Newline character  \n

- used to create a new line
	and placed within string expressions.
- when inserted in a string, all the characters
	after \n are added to a new line.

*/
