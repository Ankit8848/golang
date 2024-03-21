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
