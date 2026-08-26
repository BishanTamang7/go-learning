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

	// 3. Printf() means Print Formatted.
	// Unlike Print() and Println(), Printf() lets you control how values are displayed using format verbs.
	// 1. Basic Example
	name := "Bishan"
	age := 22

	fmt.Printf("My name is %s. I am %d Years Old.", name, age) // Output: My name is Bishan. I am 22 Years Old.

	// Here:
	// . %s -> string
	// . %d -> integer
	// . name -> replace %s
	// . age -> replace %d

	// 2. Why do we need %s and %d?
	// With Printf(), you tell Go where and how to put each value.
	// For example:
	fmt.Printf("Hello %s", name) // Output: Hello Bishan

	// %s is a format verb.
	// It tells Printf():
	// "Put a string value here"

	// 3. Most important format verbs
	// As a beginner, focus on these first:

	/*
		Verb			Used For			Example
		%s				String				"Bishan"
		%d				Integer				22
		%f				Floating-Point		3.14
		%t				Boolean				true
		%c				Character			'A'
		%v				Default value		Any Value
	*/

	// Example
	name1 := "Ram"
	age7 := 22
	height := 5.7
	isStuednt := true
	grade := 'A'
	intro := "I am Bishan Tamang"

	fmt.Printf("Name: %s", name1)
	fmt.Printf("Age: %d", age7)
	fmt.Printf("Height: %f", height)
	fmt.Printf("isStudent: %t", isStuednt)
	fmt.Printf("Grade: %c", grade)
	fmt.Printf("Introduction: %v", intro)

	// 4. %v — very useful
	// %v means:
	// Print the value in its default format.
	// Example:
	name2 := "Bishan Tamang"
	age9 := 22

	fmt.Printf("Name: %v", name2)
	fmt.Printf("Age: %v", age9)

	// You can also do:
	fmt.Printf("Name: %v, Age: %v", name2, age9)
	// For now, %v is a very useful general-purpose verb, while %s, %d, %f, etc. give you more control.

	// 5. Printf() does NOT automatically add a newline
	fmt.Printf("Hello")
	fmt.Printf("World")

	// If you want a new line, use \n:
	fmt.Printf("Hello\n")
	fmt.Printf("World\n")
	// So, Printf() does not automatically add a newline.
}
