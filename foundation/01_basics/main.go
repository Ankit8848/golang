package main

/* package main == project == workspace

Types of Packages

a. Executable => Generates a file that we can run

eg, package main -> defines a package that can be compiled and then "executed".
				Must have a func called 'main'.


b. Reusable => Code used as 'helpers'. Good place to put reusable logic

eg, package ankit -> defines a package that can be used as a dependency code(helper code).

*/

import "fmt"

/*
	import 'packages'

	eg, standard library => fmt, bufio, math, io, debug,..

	eg, reusable packages => other developer made

*/

// comments

/*

This is multiline comment.
*/

// Go tool can run file directly.
// Executables are different for OS.

// Object-Oriented Programming: Yes or NO, because it has structs.

// For Production ready : first step to take, in integrated terminal: go mod init "specify name"

/* Go CLI

go build => compiles a bunch of go source code files

go run => compiles and executes one or two files

go fmt => formats all the code in each file in the current dir

go install => compiles and "installs" a package

go get => downloads the raw source code of someone else's package

go test => runs any tests associated with the current project

*/

/*

	file organization:
		a) package main => package declaration
		b) import "fmt" => import other packages that we need
		c) func main(){
			fmt.Println("Hey there")
		}				=> declare functions, tell go to do things.
*/

func main() {
	fmt.Println("Hey everyone")
}
