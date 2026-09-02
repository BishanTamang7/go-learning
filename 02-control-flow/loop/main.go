package main

import "fmt"

func main() {
	// 1. What is loop?
	// A loop is a programming construct that allows you to execute a block of code repeatedly based on a condition.
	// It helps in automating repetitive tasks and can be used to iterate over collections, perform calculations, or control the flow of a program.

	// Without a loop:
	fmt.Println("Hello, World!")
	fmt.Println("Hello, World!")
	fmt.Println("Hello, World!")
	fmt.Println("Hello, World!")
	fmt.Println("Hello, World!")

	// With a loop:
	for i := 1; i <= 5; i++ {
		fmt.Println("Hello, World!")
	}

	// Output:
	// Hello, World!
	// Hello, World!
	// Hello, World!
	// Hello, World!
	// Hello, World!

	// So, using a loop allows us to write cleaner and more efficient code by avoiding repetition.

	// 2. Go has only one type of loop, which is the "for" loop. However, it can be used in different ways to achieve various looping behaviors.
	// Unlike some other programming languages that have multiple loop constructs (like while, do-while, etc.), Go simplifies the syntax by providing a single loop construct that can be adapted to different scenarios.

	// 3. Basic for loop syntax:
	// The basic syntax of a for loop in Go consists of three components: initialization, condition, and update statement.
	// for initialization; condition; update {
	//    // Code to be executed in each iteration
	// }

	// The initialization is executed once before the loop starts, the condition is checked before each iteration, and the update statement is executed after each iteration.

	// Example:
	for j := 0; j < 5; j++ {
		fmt.Println("Basic for loop:", j)
	}

	// Output:
	// Basic for loop: 0
	// Basic for loop: 1
	// Basic for loop: 2
	// Basic for loop: 3
	// Basic for loop: 4

	// 4. The three parts of a for loop in Go are:
	// Look at:
	// for i := 1; i <= 5; i++ {

	// There are three parts:
	// 	for initialization; condition; update
	//     		↓              ↓         ↓
	//    		start         check      change

	// Part 1: Initialization
	// i := 1
	// This happens only once at the beginning of the loop. It is used to declare and initialize the loop control variable.
	// It means that the variable i is set to 1 before the loop starts executing.

	// Part 2: Condition
	// i <= 5
	// This is checked before every iteration.
	// It means:
	//	Is i less than or equal to 5?
	// If ture -> run the loop.
	// If false -> exit the loop.

	// Part 3: Update
	// i++
	// This happens at the end of every iteration.
	// It means that the value of i is increased by 1 after each iteration of the loop.
	// 	i = 1
	//  ↓
	// 1 <= 5? → Yes → print 1
	//  ↓
	// i++ → 2
	//  ↓
	// 2 <= 5? → Yes → print 2
	//  ↓
	// i++ → 3
	//  ↓
	// 3 <= 5? → Yes → print 3
	//  ↓
	// ...
	//  ↓
	// 5 <= 5? → Yes → print 5
	//  ↓
	// i++ → 6
	//  ↓
	// 6 <= 5? → No
	//  ↓
	// STOP

	// 5. Summary of the three parts of a for loop:
	// a. Initialization: This is where you declare and initialize the loop control variable. It is executed only once at the beginning of the loop.

	// b. Condition: This is a boolean expression that is evaluated before each iteration. If the condition evaluates to true, the loop body is executed; if false, the loop terminates.
	// c. Update: This is where you modify the loop control variable after each iteration. It is executed at the end of each loop iteration.

	// 6. i++ is a shorthand for incrementing the value of i by 1. It is equivalent to writing i = i + 1. This is commonly used in loops to move to the next iteration.

	// You'll see this constantly in loops:
	// It means:
	// i = i + 1

	// For example:
	// 	i := 1
	// i++

	// Now:
	// i = 2
	// Again:
	// i++
	// Now:
	// i = 3
	// And so on...

	// 7. Counting from 0
	// You don't have to start counting from 1. You can start from 0 or any other number. For example, if you want to count from 0 to 4, you can do it like this:
	for q := 0; q < 5; q++ {
		fmt.Println("Counting from 0:", q)
	}

	// Output:
	// Counting from 0: 0
	// Counting from 0: 1
	// Counting from 0: 2
	// Counting from 0: 3
	// Counting from 0: 4

	// Notice the condtion:
	// q < 5
	// It means:
	// Is q less than 5?
	// If true → run the loop.
	// If false → exit the loop.

	// 8. Counting backwards
	// You can also count backwards by adjusting the initialization, condition, and update statements. For example, to count from 5 down to 1, you can do it like this:
	for r := 5; r > 0; r-- {
		fmt.Println("Counting backwards:", r)
	}

	// Output:
	// Counting backwards: 5
	// Counting backwards: 4
	// Counting backwards: 3
	// Counting backwards: 2
	// Counting backwards: 1

	// Here:
	// r--
	// It means:
	// r = r - 1

	// 9. Increse by more then 1
	// You can also increase the loop control variable by more than 1. For example, to count from 0 to 10 in steps of 2, you can do it like this:
	// s += 2
	for s := 0; s <= 10; s += 2 {
		fmt.Println("Increasing by more than 1:", s)
	}

	// Output:
	// Increasing by more than 1: 0
	// Increasing by more than 1: 2
	// Increasing by more than 1: 4
	// Increasing by more than 1: 6
	// Increasing by more than 1: 8
	// Increasing by more than 1: 10

	// Here i increases by 2 each time.

	// 10. You can use different variable names
	// Most programmers use i, j, k, etc. as loop control variables, but you can use any valid variable name. For example:
	for count := 1; count <= 5; count++ {
		fmt.Println("Using different variable names:", count)
	}

	// Output:
	// Using different variable names: 1
	// Using different variable names: 2
	// Using different variable names: 3
	// Using different variable names: 4
	// Using different variable names: 5

	// 11. for loop with a condition only
	// Go also allows this form:
	// for condition {
	//    // Code to be executed in each iteration
	// }
	// In this case, you can omit the initialization and update statements. The loop will continue to execute as long as the condition evaluates to true.

	// Example:
	counter := 1
	for counter <= 5 {
		fmt.Println("For loop with a condition only:", counter)
		counter++
	}

	// Output:
	// For loop with a condition only: 1
	// For loop with a condition only: 2
	// For loop with a condition only: 3
	// For loop with a condition only: 4
	// For loop with a condition only: 5

	// Here, the loop continues as long as counter is less than or equal to 5. The counter variable is incremented inside the loop body.
	// This is similar to a while loop in other programming languages.

	// Notice that:
	// counter++
	// is inside the loop body, which means it will be executed in each iteration, ensuring that the loop eventually terminates when the condition becomes false.

	// 12. Infinite loop
	// An infinite loop is a loop that continues to execute indefinitely because the condition never becomes false. In Go, you can create an infinite loop using the for statement without any conditions:
	// for {
	//    // Code to be executed indefinitely
	// }

	// Example:
	// Uncomment the following lines to see an infinite loop in action. Be cautious, as it will run indefinitely until you manually stop the program.
	/*
		for {
			fmt.Println("This is an infinite loop.")
		}
	*/

	// To exit an infinite loop, you can use a break statement or a condition that eventually becomes false.

	// This has no condition, so it keeps running.
	// Conceptually:
	// 	for
	//  ↓
	// run
	//  ↓
	// run
	//  ↓
	// run
	//  ↓
	// run
	//  ↓
	// ...

	// This is called an infinite loop.
	// For now, understand what it means, but don't use it in your practice unless you know how to stop it.

	// 13. for loop with user input
	// You can also use a for loop to repeatedly prompt the user for input until a certain condition is met. For example, you can ask the user to enter a number and continue asking until they enter a specific value.
	// Example:
	var userInput int
	for {
		fmt.Print("Enter a number (0 to exit): ")
		fmt.Scan(&userInput)
		if userInput == 0 {
			break // Exit the loop if the user enters 0
		}
		fmt.Println("You entered:", userInput)
	}

	// Output:
	// Enter a number (0 to exit): 5
	// You entered: 5
	// Enter a number (0 to exit): 10
	// You entered: 10
	// Enter a number (0 to exit): 0
}
