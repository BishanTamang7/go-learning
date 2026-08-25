package main

import "fmt"

func main() {
	// Go has three functions to output text:
	// 1. Print(). The Print() function prints its arguments with their default format. Prints everythings on the same line. Does not automatically add space.
	// Does not automatically add a newline(\n)
	var i, j string = "Hello", "Wrold"

	fmt.Print(i)
	fmt.Print(j)

	// If we want to print the arguments in new lines, we need to use \n.
	var k, l string = "Hello", "World"

	fmt.Print(k, "\n")
	fmt.Print(l, "\n")

	// It is also possible to only use one Print() for printing multiple variables
	var m, n string = "Bishan", "Tamang"

	fmt.Print(m, "\n", n)

	// If we want to add a space between string arguments, we need to use " ":
	var o, p string = "Ram", "Hari"

	fmt.Print(o, " ", p)

	// Print() inserts a space between the arguments if neither are strings:
	var q, r int = 10, 7

	fmt.Print(q, r)

	// 2. Println(). The Println() function is similar to Print() with the difference that a whitespace is added between the arguments, and a newline is added at the end:
	var s, t string = "Hello", "World"

	fmt.Println(s, t)

	// 3. Printf()
}
