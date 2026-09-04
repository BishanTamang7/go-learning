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
}
