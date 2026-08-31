package main

import "fmt"

func main() {
	// 1. What is if-else?
	// if-else is used to make a simple decision:
	//	If a condition is true, do one thing. Otherwise, do another thing.

	// For example:
	age := 22

	if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Not Adult")
	}

	// Here Go asks:
	// Is age >= 18?
	// Since 20 >= 18 is true:
	// So You are adult. is Printed.

	// 2. Basic Syntax
	// if condition {
	// 	// runs when condtion is true
	// } else {
	// 	// runs when condtion is false
	// }

	// There are always two possible paths:
	//      condition
	//      /       \
	//   true       false
	//    ↓           ↓
	// if block    else block

	// 3. Simple example

	age1 := 16

	if age1 >= 18 {
		fmt.Println("You can enter")
	} else {
		fmt.Println("You can't enter")
	}

	// Output: You can't enter
	// Because: 15 >= 18 is false.
	// So the else block runs.

	// 4. Another example
	number := 10

	if number > 5 {
		fmt.Println("Number is greater then 5")
	} else {
		fmt.Println("Number is smaller then 5")
	}

	// Output: Number is grater then 5

	// 5. if-else with equality
	number1 := 7

	if number1 == 7 {
		fmt.Println("Number is 7")
	} else {
		fmt.Println("Number isn't 7")
	}

	// Output: Number is 7

	// Remember: == means compare
	// while: = means assign.

	// So: number1 = 7 means put 7 into number1 but number1 == 7 means is number7 equal to 7?

	// 6. if-else with user input
	// This is very important example:
	var age2 int

	fmt.Print("Enter Your Age: ")
	fmt.Scan(&age2)

	if age2 >= 18 {
		fmt.Println("Yor are an adult")
	} else {
		fmt.Println("You are not an adult")
	}

	// If the user enters: 20
	// Output: You are an adult

	// if the user enters: 15
	// Output: You are not an adult

	// 7. Important rule
	// The else belongs to the if.
	// Correc:
	if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Not adult")
	}

	// Notice:
	// } else {

	// In Go, write else on the same line as the closing } of if.

	// 8. Only one of the two blocks runs.
	// if condition {
	// // this runs
	// } else {
	// // this runs
	// }

	// If condition is true:
	// if → runs
	// else → doesn't run

	// If condition is false:
	// if → doesn't run
	// else → runs
}
