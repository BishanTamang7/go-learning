package main

import "fmt"

func main() {

	// Declaring a variable
	var age int
	// var = keyword to declare a variable
	// age = variable name
	// int = data type of the variable

	// At this point, the variable 'age' has been declared but not initialized, so it holds the default value for its type, which is 0 for integers.
	fmt.Println(age)
	// Output: 0

	// Practice
	var name string
	var isStudent bool
	var height float64
	var age1 int

	fmt.Println(name)      // Output: (empty string)
	fmt.Println(isStudent) // Output: false
	fmt.Println(height)    // Output: 0
	fmt.Println(age1)      // Output: 0
}
