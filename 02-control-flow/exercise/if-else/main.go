package main

import "fmt"

func main() {
	// 1. Create a number and use if-else to check whether it is positive or negative.
	var number int

	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	if number > 0 {
		fmt.Println(number, "is Postive Number.")
	} else {
		fmt.Println(number, "is Negative Number.")
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

	if number2/5 == 0 {
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
	var password int

	fmt.Print("Enter Your Password: ")
	fmt.Scan(&password)

	if password == 12345 {
		fmt.Println("Correct Password.")
	} else {
		fmt.Println("Wrong Password.")
	}

	// 11. Take two numbers from the user and print the larger number.
	var num6 int
	var num7 int

	fmt.Print("Enter num6: ")
	fmt.Scan(&num6)
	fmt.Print("Enter num4: ")
	fmt.Scan(&num7)

	if num6 > num7 {
		fmt.Println(a, "is greater than", b)
	} else if num6 < num7 {
		fmt.Println(b, "is grater then", a)
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
}
