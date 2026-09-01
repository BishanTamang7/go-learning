package main

import "fmt"

func main() {
	// 1. What is switch?
	// A switch is used when you want to compare one value against several possible values.

	// For example, imagine you have a day:
	// 1 → Monday
	// 2 → Tuesday
	// 3 → Wednesday
	// 4 → Thursday
	// 5 → Friday
	// 6 → Saturday
	// 7 → Sunday

	// You could use many if-else statements, but switch makes this much cleaner.

	// 2. Basic switch syntax
	// swtich value {
	// 	case value1;
	// 		// code
	// 	case value2;
	// 		// code
	// 	case value3;
	// 		// code
	// 	default:
	// 		// code
	// }

	// Think of it like:
	// 	Check the value
	//      ↓
	//  ┌─────────┐
	//  │ switch  │
	//  └────┬────┘
	//       ↓
	//    case 1?
	//       ↓
	//    case 2?
	//       ↓
	//    case 3?
	//       ↓
	//    default

	// 3. Your first switch
	day := 2

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	}

	// Output: Tuesday

	// What happened?
	// Go takes:
	// day := 2
	// Then checks:
	// switch day

	// It compares day with each case:
	// 2 == 1 → false
	// 2 == 2 → true
	// It executes the code in that case:
	// fmt.Println("Tuesday")

	// 4. case
	// A case is a possible value that the switch statement can match against.
	// switch day {
	// case 1:
	// 	fmt.Println("Monday")
	// case 2:
	// 	fmt.Println("Tuesday")
	// case 3:
	// 	fmt.Println("Wednesday")
	// }

	// You can think of it as:
	// case 1 → if day == 1
	// case 2 → if day == 2
	// case 3 → if day == 3

	// 5. default
	// A default case is executed if no other case matches.
	day1 := 8
	switch day1 {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Not a valid day")
	}

	// Output: Not a valid day
	// Because:
	// 8 == 1 → false
	// 8 == 2 → false
	// 8 == 3 → false
	// None of the cases matched, so the default case was executed.

	// You can think of it as:
	// default → else

	// 6. switch with strings
	// You can also use switch with strings. Not only for numbers.
	fruit := "banana"
	switch fruit {
	case "apple":
		fmt.Println("This is an apple")
	case "banana":
		fmt.Println("This is a banana")
	case "orange":
		fmt.Println("This is an orange")
	default:
		fmt.Println("Unknown fruit")
	}

	// Output: This is a banana

	// 7. switch with user input
	// You can also use switch with user input. For example, you can ask the user to enter a number and then use switch to print the corresponding day.

	var day2 int

	fmt.Print("Enter a number (1-7): ")
	fmt.Scan(&day2)

	switch day2 {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Not a valid day")
	}

	// If the user enter 8, the output will be:
	// Not a valid day

	// 8. switch with multiple cases
	// You can also use switch with multiple cases.
	// For example, you can use switch to print the corresponding day of the week for a given number. If the number is 1, 2, or 3, it will print "Weekday".
	// If the number is 4, 5, or 6, it will print "Weekend". If the number is 7, it will print "Sunday".
	// If the number is not in the range of 1-7, it will print "Not a valid day".

	var day3 int

	fmt.Print("Enter a number (1-7): ")
	fmt.Scan(&day3)

	switch day3 {
	case 1, 2, 3:
		fmt.Println("Weekday")
	case 4, 5, 6:
		fmt.Println("Weekend")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Not a valid day")
	}

	// Output: if the user enter 8, the output will be:
	// Not a valid day

	// 9. switch does not require a break statement
	// In Go, switch statements do not require a break statement at the end of each case.
	// Once a case is matched, the code in that case is executed and the switch statement is exited automatically.
	// This is different from other programming languages like C, C++, and Java, where you need to use a break statement to exit the switch statement.

	// For example
	// switch day {
	// case 1:
	// 	fmt.Println("Monday")
	// case 2:
	// 	fmt.Println("Tuesday")
	// case 3:
	// 	fmt.Println("Wednesday")
	// }

	// If day is 2, Go prints:
	// Tuesday

	// It doesn't automatically continue into case 3.
	// So you normally don't write:
	// break inside each case, because Go automatically breaks out of the switch statement after executing the matched case.

	// 10. Switch with expressions
	// You can also use switch with expressions. For example, you can use switch to print the corresponding day of the week for a given number. If the number is 1, 2, or 3, it will print "Weekday". If the number is 4, 5, or 6, it will print "Weekend". If the number is 7, it will print "Sunday". If the number is not in the range of 1-7, it will print "Not a valid day".

	var day4 int

	fmt.Print("Enter a number (1-7): ")
	fmt.Scan(&day4)

	switch {
	case day4 >= 1 && day4 <= 3:
		fmt.Println("Weekday")
	case day4 >= 4 && day4 <= 6:
		fmt.Println("Weekend")
	case day4 == 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Not a valid day")
	}

	// Output: if the user enter 8, the output will be:
	// Not a valid day

	// 11. Summary
	// A switch statement is used to compare one value against several possible values.
	// It is a cleaner way to write multiple if-else statements.
	// It does not require a break statement at the end of each case.
	// It can be used with numbers, strings, and expressions.
}
