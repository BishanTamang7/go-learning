package main

import "fmt"

func main() {
	// 1. What is continue?
	// continue is used inside a loop to skip the current iteration and move to the next iteration.

	// The easiest way to remember:
	// break    → STOP the loop completely
	// continue → SKIP this iteration

	// 2. Basic example
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}

		fmt.Println(i)
	}

	// Output:
	// 1
	// 2
	// 4
	// 5

	// Why didn't 3 print?
	// When i becomes 3:
	// if i == 3 {
	// 	continue
	// }

	// continue says: "Don't execute the rest of this iteration. Go to the next iteration."
	// So:
	// i = 1 → print 1
	// i = 2 → print 2
	// i = 3 → continue → skip
	// i = 4 → print 4
	// i = 5 → print 5

	// 3. continue does NOT stop the loop
	// This is the biggest difference from break.

	// break
	// for i := 1; i <= 5; i++ {
	// 	if i == 3 {
	//     	break
	// 	}

	// 	fmt.Println(i)
	// }

	// Output:
	// 1
	// 2

	// The loop stops at 3.

	// continue
	// 	for i := 1; i <= 5; i++ {
	//     if i == 3 {
	//         continue
	//     }

	//	    fmt.Println(i)
	//	}

	// Output:
	// 1
	// 2
	// 4
	// 5

	// The loop continues after skipping 3.

	// 4. How continue works
	// Look at:
	// for i := 1; i <= 5; i++ {
	// 	if i == 3 {
	//     	continue
	// 	}

	// 	fmt.Println(i)
	// }

	// The flow is:
	// 	       i = 1
	//          ↓
	//      i == 3?
	//        /   \
	//      No    Yes
	//      ↓      ↓
	//   print 1  continue
	//              ↓
	//           next loop
	//              ↓
	//           i = 4

	// For i == 3, everything after continue in that iteration is skipped.

	// 5. Code after continue is skipped
	// Example:
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}

		fmt.Println("Number:", i)
		fmt.Println("Hello")
	}

	// Output:
	// Number: 1
	// Hello
	// Number: 2
	// Hello
	// Number: 4
	// Hello
	// Number: 5
	// Hello

	// When i == 3, both of these are skipped:
	// fmt.Println("Number:", i)
	// fmt.Println("Hello")

	// 6. A useful example: skip even numbers
	// You can use continue to skip values you don't want.
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println(i)
	}

	// Output:
	// 1
	// 3
	// 5
	// 7
	// 9

	// Here:
	// i%2 == 0
	//checks whether the number is even.
	// When it is even:
	// continue
	// skips that iteration.

	// 7. Skip odd numbers
	// You can do the opposite:
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			continue
		}

		fmt.Println(i)
	}

	// Output:
	// 2
	// 4
	// 6
	// 8
	// 10

	// 8. continue with user input
	for i := 1; i <= 5; i++ {
		var number int

		fmt.Print("Enter a number: ")
		fmt.Scan(&number)

		if number < 0 {
			continue
		}

		fmt.Println("You entered:", number)
	}

	// If the user enters a negative number, the program skips:
	// fmt.Println("You entered:", number)
	// and moves to the next loop iteration.

	// 9. continue vs break
	// This is the most important thing to understand.

	// break
	for i := 1; i <= 5; i++ {
		if i == 3 {
			break
		}

		fmt.Println(i)
	}

	// Output: 1, 2
	// Meaning: Stop the entire loop.

	// continue
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}

		fmt.Println(i)
	}
	// Output: 1, 2, 3, 4, 5
	// Meaning: Skip only the current iteration and keep looping.

	// 10. Easy mental model
	// Imagine you're checking students one by one.

	// Student 1 → process
	// Student 2 → process
	// Student 3 → skip → continue
	// Student 4 → process
	// Student 5 → process

	// continue means: Skip this one, but keep going."
	// break means: "Stop checking everyone."

	// 11. continue with a condition
	// The most common pattern is:
	// for ... {
	// 	if condition {
	//     	continue
	// 	}

	// 	// code
	// }

	// For example:
	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue
		}

		fmt.Println(i)
	}

	// Output: 1, 2, 3, 4, 6, 7, 8, 9, 10

	// 12. Important: continue works with loops
	// You generally use it inside a for loop:
	// for condition {
	// 	if condition {
	//     	continue
	// 	}

	// 	// code
	// }

	// It doesn't mean "skip the whole program."
	// It means: Skip the remaining code in the current loop iteration.

	// 13. Remember these two forever
	// break
	//   ↓
	// STOP LOOP

	// continue
	// 	↓
	// SKIP CURRENT ITERATION
	// 	↓
	// NEXT ITERATION
}
