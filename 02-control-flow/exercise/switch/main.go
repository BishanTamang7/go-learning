package main

import "fmt"

func main() {
	// 1. Create day := 1 and print:
	// 1 → Monday
	// 2 → Tuesday
	// 3 → Wednesday
	// ...
	// 7 → Sunday

	day := 1

	switch day {
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
	}

	// 2. Number to Word
	// 1 → One
	// 2 → Two
	// 3 → Three
	// 4 → Four
	// 5 → Five

	var number int

	fmt.Print("Enter a Number: ")
	fmt.Scan(&number)

	switch number {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Invalid Number")
	}

	// 3.Create a month number and print its month name.
	var month int

	fmt.Print("Enter a Month Number: ")
	fmt.Scan(&month)

	switch month {
	case 1:
		fmt.Println("January")
	case 2:
		fmt.Println("February")
	case 3:
		fmt.Println("March")
	case 4:
		fmt.Println("April")
	case 5:
		fmt.Println("May")
	case 6:
		fmt.Println("June")
	case 7:
		fmt.Println("July")
	case 8:
		fmt.Println("August")
	case 9:
		fmt.Println("September")
	case 10:
		fmt.Println("October")
	case 11:
		fmt.Println("November")
	case 12:
		fmt.Println("December")
	default:
		fmt.Println("Invalid Number")
	}

	// 4. Menu Choice
	// 1 → Start
	// 2 → Settings
	// 3 → Help
	// 4 → Exit

	var menu int

	fmt.Print("Enter a Menu: ")
	fmt.Scan(&menu)

	switch menu {
	case 1:
		fmt.Println("Start")
	case 2:
		fmt.Println("Settings")
	case 3:
		fmt.Println("Help")
	case 4:
		fmt.Println("Exit")
	}
}
