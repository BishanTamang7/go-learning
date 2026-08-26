package main

import "fmt"

func main() {
	// User input in Go means allowing the person running your program to enter some information.
	// 1. Taking one input with fmt.scan()
	var name string

	fmt.Print("Enter Your Name: ")
	fmt.Scan(&name)

	fmt.Println("Hello", name)

	// What is happening?
	// Let's break it down:

	// var name string -> Creates a variable called name.
	// fmt.Print("Enter Your Name: ") -> display Enter Your Name on Consol.
	// fmt.Scan(&name) -> waits for the user to enter somethings.
	// If you enter: Bishan then go stores it inside name variable and so now like this name = "Bishan".
	// fmt.Println("Hello", name)-> it prints Hello Bishan

	// 2. Why do we use &name?
	// This is one of the most important things to understand.
	// We write:
	// fmt.Scan(&name)
	// Not:
	// fmt.Scan(name)
	// The & means:
	// Give Scan() the location of the variable so it can put the user's input there.

	// 3. You can also use short variable declaration?
	// This connect directly to what you asked about previously.
	// You can't use := directly with fmt.Sacn() like this:
	// name := fmt.Scan()
	// x This does not work the way you want.

	// Why?
	// Because fmt.Scan() neeeds a variable where it can store the input.
	// So, normally do:
	// var name string
	// fmt.Scan(&name)
	// Or, for simple cases, you can declare the variable first and then scan into it.

	// 4. fmt.Scan() and spaces
	// This is important.
	// Suppose:
	// var name string
	// fmt.Scan(&name)

	// You Enter:
	// Bishan Tamang
	// Scan() will read only:
	// Scan() will read only:
	// because Scan() treats whitespace as a separator.
	// So:
	/*
	   Bishan Tamang
	         ↑
	        space
	*/
}
