package main

import "fmt"

func main() {
	// 1. What is break?
	// The break statement is used to exit a loop prematurely, before the loop condition is false.
	// It allows you to terminate the loop and continue executing the code after the loop.

	// Example:
	for i := 1; i <= 10; i++ {
		if i == 5 {
			break // Exit the loop when i is equal to 5
		}
		println(i)
	}

	// Output:
	// 1
	// 2
	// 3
	// 4

	// In this example, the loop will print numbers from 1 to 4, and when i becomes 5, the break statement will be executed, causing the loop to terminate.

	// The break statement can also be used in switch statements to exit a case block early.

	// 2. Basic Syntax
	// break

	// You normally use it inside a loop
	// for condition {
	// if someCondition {
	//     break
	// 	}
	// }

	// 3. How break works
	// Look at this:

	for i := 1; i <= 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println(i)
	}

	// The execution is:
	// i = 1 → i == 5? No  → print 1
	// i = 2 → i == 5? No  → print 2
	// i = 3 → i == 5? No  → print 3
	// i = 4 → i == 5? No  → print 4
	// i = 5 → i == 5? Yes → BREAK

	// The loop ends.

	// It does not continue to:
	// 6
	// 7
	// 8
	// 9
	// 10

	// 4. beark with user input
	// You can use break with user input to exit a loop based on a condition.

	var number int
	for {
		fmt.Print("Enter a number (0 to exit): ")
		fmt.Scanln(&number)

		if number == 0 {
			break // Exit the loop if the user enters 0
		}

		fmt.Println("You entered:", number)
	}

	fmt.Println("Loop exited.")

	// If the user enters: 5, 10, 15, 0
	// The output will be:
	// You entered: 5
	// You entered: 10
	// You entered: 15
	// Loop exited.

	// When the user enters 0:
	// 	if number == 0 {
	//     break
	// }

	// the loop stop.

	// 5. break with a counting loop
	for i := 1; i <= 10; i++ {
		if i == 5 {
			break // Exit the loop when i is equal to 5
		}
		fmt.Println(i)
	}

	// Output:
	// 1
	// 2
	// 3
	// 4

	// Even though the loop says:
	// i <= 100
	// break stops it when i reaches 10.

	// 6. Why would we use break?
	// Sometimes you don't know exactly when you want the loop to stop.

	// For example, imagine searching for a number:
	// numbers := []int{10, 20, 30, 40, 50}
	// You want to stop searching as soon as you find 30.

	// Conceptually:
	// 	Search 10 → not found
	// 	Search 20 → not found
	// 	Search 30 → FOUND
	//              ↓
	//            break
	//              ↓
	//         stop searching

	// We'll work more with slices and arrays later. For now, understand the idea.

	// 7. break with if
	// The most common pattern is:
	// for ... {
	// 	if condition {
	//     	break
	// 		}
	// }

	// For example:
	// for i := 1; i <= 20; i++ {
	// 	if i > 5 {
	// 		break
	// 	}

	// 	fmt.Println(i)
	// }

	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5

	// The moment:
	// i > 5
	// becomes true, the loop stops.

	// 8. break vs normal loop condition
	// Compare these:

	// Without break
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// The loop naturally stops when: i <= 5 become false.

	// With break
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}

	// The loop stops when: i > 5 becomes true, and the break statement is executed.

	// 9. break in an infinite loop
	// Remember the infinite loop we saw earlier?
	// for {
	// 	fmt.Println("Hello")
	// }

	// It never stops by itself.

	// You can use break to give it an exit condition:

	for {
		var number int

		fmt.Print("Enter a number: ")
		fmt.Scan(&number)

		if number == 0 {
			break
		}

		fmt.Println(number)
	}

	// Now:
	// 	Enter 5 → continue
	// Enter 8 → continue
	// Enter 2 → continue
	// Enter 0 → break → stop

	// This is a very common pattern.

	// 10. break inside nested loops
	// You can have a loop inside another loop:
	for i := 1; i <= 3; i++ {

		for j := 1; j <= 3; j++ {

			if j == 2 {
				break
			}

			fmt.Println(i, j)
		}
	}

	// Here, the break stops the inner loop.
	// It doesn't automatically stop the outer loop.
	// For now, don't worry too much about nested loops. Just remember:
	// A normal break stops the loop it is currently inside.

	// 11. Important rule
	// break is mainly used with loops.

	// Think:
	// 	for
	//  ↓
	// repeat
	//  ↓
	// condition becomes true
	//  ↓
	// break
	//  ↓
	// STOP LOOP

	// 12. The key difference

	// for
	// "keep repeating"
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	// break
	// "Stop repeating now."
	for i := 1; i <= 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println(i)
	}

	// So remember:
	// for   → repeat
	// break → stop the loop
}
