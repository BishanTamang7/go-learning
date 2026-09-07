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
	default:
		fmt.Println("Invalid menu choice.")
	}

	// 5. Traffic Light
	// "red" → Stop
	// "yellow" → Wait
	// "green" → Go
	var trafficLight string

	fmt.Print("Enter a Light: ")
	fmt.Scan(&trafficLight)

	switch trafficLight {
	case "red":
		fmt.Println("Stop")
	case "yellow":
		fmt.Println("Wait")
	case "green":
		fmt.Println("Go")
	default:
		fmt.Println("Invalid traffic light")
	}

	// 6. Grade. Create a grade variable:
	// "A" → Excellent
	// "B" → Very Good
	// "C" → Good
	// "D" → Pass
	// "E" → Fail

	var grade rune

	fmt.Print("Enter a grade: ")
	fmt.Scan(&grade)

	switch grade {
	case 'A':
		fmt.Println("Excellent")
	case 'B':
		fmt.Println("Very Good")
	case 'C':
		fmt.Println("Good")
	case 'D':
		fmt.Println("Pass")
	case 'E':
		fmt.Println("Fail")
	default:
		fmt.Println("Invalid Grade.")
	}

	// 7. Calculator Operator. Create an operator:
	// + → Addition
	// - → Subtraction
	// * → Multiplication
	// / → Division

	var operator string

	fmt.Print("Enter an Operator: ")
	fmt.Scan(&operator)

	switch operator {
	case "+":
		fmt.Println("Addition")
	case "-":
		fmt.Println("Subtraction")
	case "*":
		fmt.Println("Multiplication")
	case "/":
		fmt.Println("Division")
	default:
		fmt.Println("Invalid Operator.")
	}

	// 8. Season
	// 1 → Spring
	// 2 → Summer
	// 3 → Autumn
	// 4 → Winter

	var season int

	fmt.Print("Enter a season: ")
	fmt.Scan(&season)

	switch season {
	case 1:
		fmt.Println("Spring")
	case 2:
		fmt.Println("Summer")
	case 3:
		fmt.Println("Autumn")
	case 4:
		fmt.Println("Winter")
	default:
		fmt.Println("Invalid number.")
	}

	// 9. Simple Menu
	// 1 → View Profile
	// 2 → Edit Profile
	// 3 → Logout

	var simpleMenu int

	fmt.Print("Enter a simple menu: ")
	fmt.Scan(&simpleMenu)

	switch simpleMenu {
	case 1:
		fmt.Println("View Profile")
	case 2:
		fmt.Println("Edit Profile")
	case 3:
		fmt.Println("Logout")
	default:
		fmt.Println("Invalid number")
	}

	// 10. Language Selection
	// 1 → English
	// 2 → Nepali
	// 3 → Hindi
	// 4 → Japanese

	var language int

	fmt.Print("Enter a Language: ")
	fmt.Scan(&language)

	switch language {
	case 1:
		fmt.Println("English")
	case 2:
		fmt.Println("Nepali")
	case 3:
		fmt.Println("Hindi")
	case 4:
		fmt.Println("Japanese")
	default:
		fmt.Println("Invalid number.")
	}

	// 11. Create a day number and use one case for weekdays and another for weekends.
	var day7 int

	fmt.Print("Enter a day: ")
	fmt.Scan(&day7)

	switch day7 {
	case 1, 2, 3, 4, 5:
		fmt.Println("Weekdays")
	case 6, 7:
		fmt.Println("Weekends")
	default:
		fmt.Println("Invalid day number.")
	}

	// 12. Month Days
	// Take a month number and print the number of days for months where the answer is straightforward. Handle invalid numbers with default.

	var monthDay int

	fmt.Print("Enter  month number (1-12): ")
	fmt.Scan(&monthDay)

	switch monthDay {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println("31 days")
	case 4, 6, 9, 11:
		fmt.Println("30 days")
	case 2:
		fmt.Println("28 or 29 days")
	default:
		fmt.Println("Invalid month number")
	}

	// 13. Character Type
	// 'a' → Vowel
	// 'e' → Vowel
	// 'i' → Vowel
	// 'o' → Vowel
	// 'u' → Vowel

	var characterType rune

	fmt.Print("Enter a character type (a,e,i,o,u): ")
	fmt.Scan(&characterType)

	switch characterType {
	case 'a', 'e', 'i', 'o', 'u':
		fmt.Println("Vowel")
	default:
		fmt.Println("Consonant")
	}

	// 14. Simple ATM Menu
	// 1 → Check Balance
	// 2 → Deposit
	// 3 → Withdraw
	// 4 → Exit

	var atm int

	fmt.Print("Enter an ATM Menu (1-4): ")
	fmt.Scan(&atm)

	switch atm {
	case 1:
		fmt.Println("Check Balance")
	case 2:
		fmt.Println("Deposit")
	case 3:
		fmt.Println("Withdraw")
	case 4:
		fmt.Println("Exit")
	default:
		fmt.Println("Invalid ATM Menu.")
	}

	// 15. Food Menu
	// 1 → Pizza
	// 2 → Burger
	// 3 → Momo
	// 4 → Chowmein
	// 5 → Exit

	var foodMenu int

	fmt.Print("Enter a food menu (1-5): ")
	fmt.Scan(&foodMenu)

	switch foodMenu {
	case 1:
		fmt.Println("Pizza")
	case 2:
		fmt.Println("Burger")
	case 3:
		fmt.Println("Momo")
	case 4:
		fmt.Println("Chowmein")
	case 5:
		fmt.Println("Exit")
	default:
		fmt.Println("Invalid food menu.")
	}

	// 16. HTTP Status Code
	// 200 → OK
	// 404 → Not Found
	// 500 → Server Error

	var httpStatus int

	fmt.Print("Enter a HTTP Status Code: ")
	fmt.Scan(&httpStatus)

	switch httpStatus {
	case 200:
		fmt.Println("OK")
	case 404:
		fmt.Println("Not Found")
	case 500:
		fmt.Println("Server Error")
	default:
		fmt.Println("Invalid HTTP Status Code.")
	}

	// 17. Number Category. Use switch to handle:
	// 1 → Beginner
	// 2 → Intermediate
	// 3 → Advanced

	var numberCategory int

	fmt.Print("Enter a Number Category (1-3): ")
	fmt.Scan(&numberCategory)

	switch numberCategory {
	case 1:
		fmt.Println("Beginner")
	case 2:
		fmt.Println("Intermediate")
	case 3:
		fmt.Println("Advanced")
	default:
		fmt.Println("Invalid Number Category.")
	}

	// 18. File Extension. Create an extension:
	// "jpg" → Image
	// "png" → Image
	// "mp3" → Audio
	// "mp4" → Video
	// "pdf" → Document

	var fileExtension string

	fmt.Print("Enter a File Extension: ")
	fmt.Scan(&fileExtension)

	switch fileExtension {
	case "jpg":
		fmt.Println("Image")
	case "png":
		fmt.Println("Image")
	case "mp3":
		fmt.Println("Audio")
	case "mp4":
		fmt.Println("Video")
	case "pdf":
		fmt.Println("Document")
	default:
		fmt.Println("Invalid File Extension.")
	}

	// 19. Operation Menu
	// Take a choice from the user and perform one of several operations based on the choice.
	var choice int
	var num1, num2 float64

	fmt.Println("Operation Menu")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")

	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)

	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	switch choice {
	case 1:
		fmt.Println("Result:", num1+num2)

	case 2:
		fmt.Println("Result:", num1-num2)

	case 3:
		fmt.Println("Result:", num1*num2)

	case 4:
		if num2 == 0 {
			fmt.Println("Cannot divide by zero")
		} else {
			fmt.Println("Result:", num1/num2)
		}

	default:
		fmt.Println("Invalid choice")
	}

	// 20. Calculator
	// Take two numbers and an operator from the user.
	// Use switch to perform +, -, *, /

	var a, b int
	var op string

	fmt.Print("Enter first number: ")
	fmt.Scan(&a)

	fmt.Print("Enter operator: ")
	fmt.Scan(&op)

	fmt.Print("Enter second number: ")
	fmt.Scan(&b)

	switch op {
	case "+":
		fmt.Println("Answer:", a+b)

	case "-":
		fmt.Println("Answer:", a-b)

	case "*":
		fmt.Println("Answer:", a*b)

	case "/":
		if b == 0 {
			fmt.Println("Cannot divide by zero")
		} else {
			fmt.Println("Answer:", a/b)
		}

	default:
		fmt.Println("Invalid operator")
	}
}
