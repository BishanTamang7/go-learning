package main

import "fmt"

func main() {
	/*
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
	*/

	/*
		// Declaring and initializing a variable

		// Declaring a variable
		var age int

		// Assigning a value to the variable
		age = 25

		// Remember:
		// Declaring -> Creating a variable and specifying its type.
		// Assigning -> Giving a value to the variable.

		fmt.Println(age) // Output: 25

		// Practice
		// Declaring variables
		var name string
		var isStudent bool
		var height float64
		var age1 int

		// Assigning values to the variables
		name = "Bishan Tamang"
		isStudent = true
		height = 5.7
		age1 = 22

		// Printing the values of the variables
		fmt.Println(name)      // Output: Bishan Tamang
		fmt.Println(isStudent) // Output: true
		fmt.Println(height)    // Output: 5.7
		fmt.Println(age1)      // Output: 22
	*/

	/*
		// Declaring and initializing a variable in one line
		var age int = 25

		fmt.Println(age) // Output: 25

		// Practice
		var name string = "Bishan Tamang"
		var isStudent bool = true
		var height float64 = 5.7
		var age1 int = 22

		fmt.Println(name)      // Output: Bishan Tamang
		fmt.Println(isStudent) // Output: true
		fmt.Println(height)    // Output: 5.7
		fmt.Println(age1)      // Output: 22
	*/

	/*
		// Type inference
		// Go can infer the type of a variable based on the value assigned to it.
		// This means you can declare and initialize a variable without explicitly specifying its type.

		// Declaring and initializing a variable with type inference
		var age = 25 // Go infers that 'age' is of type int

		fmt.Println(age) // Output: 25

		// Practice
		var name = "Bishan Tamang"
		var isStudent = true
		var height = 5.7
		var age1 = 22

		fmt.Println(name)      // Output: Bishan Tamang
		fmt.Println(isStudent) // Output: true
		fmt.Println(height)    // Output: 5.7
		fmt.Println(age1)      // Output: 22
	*/

	/*
		// Short variable declaration
		// Go provides a shorthand syntax for declaring and initializing variables using the := operator.
		// This is known as short variable declaration.

		// Declaring and initializing a variable using short variable declaration
		age := 25 // Go infers that 'age' is of type int

		fmt.Println(age) // Output: 25

		// practice
		name := "Bishan Tamang"
		isStudent := true
		height := 5.7
		age1 := 22

		fmt.Println(name)      // Output: Bishan Tamang
		fmt.Println(isStudent) // Output: true
		fmt.Println(height)    // Output: 5.7
		fmt.Println(age1)      // Output: 22
	*/

	// Assignment "=" vs Declaration ":="
	// The "=" operator is used for assignment, while the ":=" operator is used for declaration and initialization.
	// The ":=" operator can only be used inside functions, while the "=" operator can be used both inside and outside functions.

	// Declaring and initializing a variable using short variable declaration
	age := 25 // Go infers that 'age' is of type int

	fmt.Println(age) // Output: 25

	// Assigning a new value to the variable using the "=" operator
	age = 30

	fmt.Println(age) // Output: 30

	// I can't do this:
	// age := 35
	// age := 35
	// This will cause a compilation error because 'age' has already been declared in this scope.

	// However, I can do this:
	age = 35 // This is valid because I'm just assigning a new value to the existing variable 'age'.

	fmt.Println(age) // Output: 35

	// Remember:
	// "=" is used for assignment/change an existing variable.
	// ":=" is used for declaration and initialization.
}
