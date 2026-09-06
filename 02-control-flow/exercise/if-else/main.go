package main

import "fmt"

func main() {
	// 1. Create a number and use if-else to check whether it is positive or negative.
	var number int

	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	if number > 0 {
		fmt.Println(number, "is Positive Number.")
	} else if number < 0 {
		fmt.Println(number, "is Negative Number.")
	} else {
		fmt.Println("The number is Zero.")
	}

	// 2. Create a number and check whether it is even or odd.
	var number1 int

	fmt.Print("Enter a Number: ")
	fmt.Scan(&number1)

	if number1%2 == 0 {
		fmt.Println(number1, "is Even Number.")
	} else {
		fmt.Println(number1, "is Odd Number.")
	}

	// 3. Create an age variable. If age is 18 or above, print "Adult", otherwise print "Minor".
	var age int

	fmt.Print("Enter a Age: ")
	fmt.Scan(&age)

	if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}

	// 4. Create marks. If marks are 40 or above, print "Pass", otherwise print "Fail".
	var marks int

	fmt.Print("Enter a Marks: ")
	fmt.Scan(&marks)

	if marks >= 40 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}

	// 5. Create two numbers and print which number is greater.
	a := 6
	b := 7

	if a > b {
		fmt.Println(a, "is greater than", b)
	} else if b > a {
		fmt.Println(b, "is greater than", a)
	} else {
		fmt.Println("They are equal")
	}

	// 6. Create two numbers and check whether they are equal.
	c := 8
	d := 9

	if c == d {
		fmt.Println(c, "and", d, "are equal.")
	} else {
		fmt.Println(c, "and", d, "are not equal.")
	}

	// 7. Create a number and check whether it is divisible by 5.
	number2 := 10

	if number2%5 == 0 {
		fmt.Println(number2, "is divisible by 5.")
	} else {
		fmt.Println(number2, "is not divisible by 5.")
	}

	// 8. Create a temperature. If it is greater than 30, print "Hot", otherwise print "Not Hot".
	var temperature int

	fmt.Print("Enter a temperature: ")
	fmt.Scan(&temperature)

	if temperature > 30 {
		fmt.Println("Hot")
	} else {
		fmt.Println("Not Hot")
	}

	// 9. Create an age. If age is 18 or above, print "You can vote", otherwise print "You cannot vote".
	var age3 int

	fmt.Print("Enter Your Age: ")
	fmt.Scan(&age3)

	if age3 >= 18 {
		fmt.Println("You can vote.")
	} else {
		fmt.Println("You cannot vote.")
	}

	// 10. Create a password variable. If it equals "12345", print "Correct Password", otherwise print "Wrong Password".
	var password string

	fmt.Print("Enter Your Password: ")
	fmt.Scan(&password)

	if password == "12345" {
		fmt.Println("Correct Password.")
	} else {
		fmt.Println("Wrong Password.")
	}

	// 11. Take two numbers from the user and print the larger number.
	var num6 int
	var num7 int

	fmt.Print("Enter num6: ")
	fmt.Scan(&num6)
	fmt.Print("Enter num7: ")
	fmt.Scan(&num7)

	if num6 > num7 {
		fmt.Println(num6, "is greater than", num7)
	} else if num6 < num7 {
		fmt.Println(num7, "is grater then", num6)
	} else {
		fmt.Println("Both are equal")
	}

	// 12. Take a number. Check whether it is between 1 and 100.
	var number7 int

	fmt.Print("Enter a Number: ")
	fmt.Scan(&number7)

	if number7 >= 1 && number7 <= 100 {
		fmt.Println(number7, "it is between 1 and 100.")
	} else {
		fmt.Println(number7, "it is not between 1 and 100.")
	}

	// 13. Create an order amount. If it is 1000 or more, print "Free Delivery", otherwise print "Delivery Charge Applies".
	var order_amount int

	fmt.Print("Enter order amount: ")
	fmt.Scan(&order_amount)

	if order_amount >= 1000 {
		fmt.Println("Free Delivery.")
	} else {
		fmt.Println("Delivery Charge Applies.")
	}

	// 14. Create username and password. Check whether both match predefined values.
	var username string
	var password1 int

	fmt.Print("Enter a Username: ")
	fmt.Scan(&username)

	fmt.Print("Enter a Password: ")
	fmt.Scan(&password1)

	if username == "bishan" && password1 == 123 {
		fmt.Println("Both match predefined values.")
	} else {
		fmt.Println("Both did not match predefined values.")
	}

	// 15. Take a number from the user and check whether it is a multiple of 10.
	var number3 int

	fmt.Print("Enter a Number: ")
	fmt.Scan(&number3)

	if number3%10 == 0 {
		fmt.Println(number3, "is multiple of 10.")
	} else {
		fmt.Println(number3, "is not multiiple of 10.")
	}

	// 16. Take an age and print "Child" if below 13, otherwise "Not Child".
	var age10 int

	fmt.Print("Enter Your Age: ")
	fmt.Scan(&age10)

	if age10 < 13 {
		fmt.Println("Child.")
	} else {
		fmt.Println("Not Child.")
	}

	// 17. Create two players' scores. Print "Player 1 Wins" if the first score is greater, otherwise "Player 2 Wins".
	var scores1 int
	var scores2 int

	fmt.Print("Enter Player 1 Score: ")
	fmt.Scan(&scores1)

	fmt.Print("Enter Player 2 Score: ")
	fmt.Scan(&scores2)

	if scores1 > scores2 {
		fmt.Println("Player 1 Wins.")
	} else if scores2 > scores1 {
		fmt.Println("Player 2 Wins.")
	} else {
		fmt.Println("It's a Tie")
	}

	// 18. Take a number and check whether it is exactly 100.
	var number77 int

	fmt.Print("Enter a Number: ")
	fmt.Scan(&number77)

	if number77 == 100 {
		fmt.Println("Exactly 100.")
	} else {
		fmt.Println("Not Exectly 100.")
	}

	// 19. Create a salary. If it is greater than 50,000, print "High Salary", otherwise print "Regular Salary".
	var salary int

	fmt.Print("Enter a Salary: ")
	fmt.Scan(&salary)

	if salary > 50000 {
		fmt.Println("High Salary.")
	} else {
		fmt.Println("Regular Salary.")
	}

	// 20. Take a password string. If its length is at least 8, print "Strong Enough", otherwise print "Too Short".
	var password11 string

	fmt.Print("Enter a Password: ")
	fmt.Scan(&password11)

	if len(password11) >= 8 {
		fmt.Println("Strong Enough")
	} else {
		fmt.Println("Too Short.")
	}
}
