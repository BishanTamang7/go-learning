package main

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

	/*
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
	*/

	/*
		// Multiple variable declaration and initialization
		// Go allows you to declare and initialize multiple variables in a single line.

		// Declaring and initializing multiple variables in one line
		var name, age, isStudent = "Bishan Tamang", 22, true

		// Declaring and initializing multiple variables using short variable declaration
		height, weight := 5.7, 60.5

		// Printing the values of the variables
		println(name)      // Output: Bishan Tamang
		println(age)       // Output: 22
		println(isStudent) // Output: true
		println(height)    // Output: 5.7
		println(weight)    // Output: 60.5
	*/

	/*
		// Variable/Go Naming Conventions
		// 1. Variable names can contain letters, digits, and underscores.
		// Good
		userName := "Bishan"
		user123 := 123
		user_name := "ABC" // valid, but not idiomatic Go

		// Bad
		user - name := "Bishan" // hyphs is not allowed
		user.name := "bishan"   // dot is not allowed in a variable declaration

		// 2. Variable names must start with a letter or an underscore.
		// Good
		userName := "Bishan"
		_userName := "Bishan"

		// Bad
		123user := "Bishan" // variable names cannot start with a digit

		// 3. Variable names are case-sensitive.
		name := "Bishan"
		Name := "Tamang" // This is a different variable from 'name' because Go is case-sensitive.

		fmt.Println(name) // Output: Bishan
		fmt.Println(Name) // Output: Tamang

		// 4. Variable names should not be Go keywords.
		// Good
		var age int = 22 // 'age' is not a Go keyword

		// Bad
		var var int = 22 // 'var' is a Go keyword, so this will cause a compilation error.
		func := "myFunction" // 'func' is a Go keyword, so this will cause a compilation error.

		// 5. Go generally uses camelCase for variable names.
		// Good
		firstName := "Bishan"
		totalAmount := 100.50
		userAge := 22

		// Valid but not idiomatic Go
		First_name := "Bishan"
		total_amount := 100.50
	*/

	/*
		// Constants
		// Constants are immutable values that are known at compile time and do not change during the execution of a program.
		// They are declared using the 'const' keyword.

		// Declaring a constant
		const pi = 3.14

		// Printing the value of the constant
		println(pi) // Output: 3.14

		// Practice
		const gravity = 9.81
		const speedOfLight = 299792458 // in meters per second

		println(gravity)      // Output: 9.81
		println(speedOfLight) // Output: 299792458
	*/

	// Type Conversion
	// In Go, you can convert a value from one type to another using type conversion.
	// Type conversion is done by specifying the target type in parentheses before the value to be converted.

	// Example of type conversion
	var age int = 25
	var ageFloat float64 = float64(age) // Converting int to float64

	println(age)      // Output: 25
	println(ageFloat) // Output: 25.0

	// Go does not perform implicit type conversion, so you must explicitly convert types when necessary.
	// For example,
	var num1 int = 10
	var num2 float64 = 5.5

	// var result = num1 + num2 // This will cause a compilation error because num1 is an int and num2 is a float64.
	// To fix this, you need to convert one of the values to the other type:
	var resultFixed = float64(num1) + num2 // Now both values are of type float64, so the addition is valid.

	println(resultFixed) // Output: 15.5

	// Go also does not allow you to convert between types that are not compatible, such as converting a string to an int.
	// Go does not round numbers when converting from float to int; it simply truncates the decimal part.
	// Converting a float64 to an int will result in the decimal part being discarded.
	// If you try to convert a float64 value of 5.9 to an int, the result will be 5, not 6.

	// Practice
	var height float64 = 5.9
	var heightInt int = int(height) // Converting float64 to int

	println(height)    // Output: 5.9
	println(heightInt) // Output: 5

}
