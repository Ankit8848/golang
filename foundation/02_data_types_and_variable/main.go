package main

import "fmt"

/*  Types

- case-insensitive.
- variable type specify before executing.
- Everything is a Type.

String(16 bytes)
Bool(1 byte)	true  false
Integer 	uint8	uint64	int8	int64	uintptr 	alias
Floating	float32		float64
Complex		complex64	complex128

complex64=> 8 bytes     1+2i
complex128 => 16 bytes  4+2i

const PI float64 = 3.141592 - like enums or env. vars.


Others:  Arrays	 Slices		Maps	Structs		Pointers	Functions	Channels ...



For	Integer

	Data Type		Memory

		uint8		8 bits or 1 byte
		uint16		16 bits or 2 bytes
		uint32		32 bits or 4 bytes
		uint64		64 bits or 8 bytes
		int8		8 bits or 1 byte
		int16		16 bits or 2 bytes
		int32		32 bits or 4 bytes
		int64		64 bits or 8 bytes
		int			4 bytes for 32-bit machines, 8 bytes for 64-bit machines => alias are - rune

		-uint means unsigned integer. where only + value include.
		-int means signed integer. where both + or - value include.

For Float

	Data Type		memory

	float32			32 bits or 4 bytes
	float64			64 bits or 8 bytes


*/

// for printing  a variable and string
// 1st way to initialize variable

// Go is statically typed.
// variables are assigned a type, either explicitly or implicitly.

// syntax: var <variable name> <data type> = <Value>

const LoginToken string = "LoveYouAll"

// note: 1st character of variable is capital because we want to tell that it is public.
// so, it can be accessible anywhere.

// Public => var name starts with capital letter
// private => var name starts with small letter

func main() {
	var username string
	fmt.Println("Username declared: ", username)
	// created but not declared so, it will give empty string.

	username = "Ankit Jha"
	fmt.Println("Username assigned: ", username)
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

/*
	Zero values

	Verb		Description

	int			0
	float64		0.0
	string		""
	bool		false

	pointers,		nil
	functions,
	interfaces,maps
*/
