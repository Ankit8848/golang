package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hey there")
	fmt.Println("Please rate this repo between 1 to 10 scale")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)

	// with strings package - and trim the parsing error. which is caused due to click enter.
	numbConv, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Add 9 to your rating: ", numbConv+9)
	}

}

/*

	To know the Type of variable

	1st way: using %T - format specifier

	2nd way: using reflect.TypeOf()
			- function from the reflect package.

	ex: import "reflect"
	func main() {
		var name1 string = "ankit"
		fmt.printf("variable value= %v is of type %v \n",name1, reflect.TypeOf(name1))
		}
*/

/*
	type casting -
		- The process of converting one data type to another is known as Type casting.
		- Data types can be converted to other data types, but this does not guarantee
			that the value will remain intact.


*/
