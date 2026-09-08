package main

import "fmt"

func main() {
	// 1. Stop at 5
	// Print 1–10 but stop when the number reaches 5.
	for a := 1; a <= 10; a++ {
		if a == 5 {
			break
		}
		fmt.Println(a)
	}

	// 2. Stop at 10
	// Print 1–100 but stop when the number reaches 10.
	for b := 1; b <= 100; b++ {
		if b == 10 {
			break
		}
		fmt.Println(b)
	}

	// 3. Stop at 20
	// Print 1–100 but stop at 20.
	for c := 1; c <= 100; c++ {
		if c == 20 {
			break
		}
		fmt.Println(c)
	}

	// 4. Stop at 1
	// Count backward from 10 and use break when you reach 1.
	for d := 10; d >= 1; d-- {
		if d == 1 {
			break
		}
		fmt.Println(d)
	}

	// 5. Stop at Even Number
	// Loop from 1 upward and stop when you encounter the first even number.
	for e := 1; e <= 10; e++ {
		if e%2 == 0 {
			break
		}
		fmt.Println(e)
	}

	// 6. Stop at Odd Number
	// Loop from 1 upward and stop when you encounter the first odd number.
	for f := 1; f <= 10; f++ {
		if f%2 != 0 {
			break
		}
		fmt.Println(f)
	}

	// 7. Stop at Multiple of 10
	// Loop from 1 and stop when you reach a number divisible by 10.
	for g := 1; g <= 20; g++ {
		if g%10 == 0 {
			break
		}
		fmt.Println(g)
	}

	// 8. Stop at 50
	// Print numbers from 1 and stop when you reach 50.
	for i := 1; i <= 100; i++ {
		if i == 50 {
			break
		}
		fmt.Println(i)
	}

	// 9. Stop at User Number
	// Take a number from the user and stop the loop when the counter reaches that number.
	var number int

	fmt.Println("Enter a number: ")
	fmt.Scan(&number)

	for j := 1; j <= number; j++ {
		if j == number {
			break
		}
		fmt.Println(j)
	}

	// 10. Stop at 0
	// Continuously take numbers from the user and stop when the user enters 0.
	for {
		var num int

		fmt.Println("Enter a number: ")
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		fmt.Println(num)
	}

	// 11. Password Attempts
	// Allow the user to enter a password repeatedly. Stop the loop when the correct password is entered.
	for {
		var password int

		fmt.Print("Enter a Password: ")
		fmt.Scan(&password)

		if password == 1234 {
			break
		}
		fmt.Println("Incorrect Password")
	}
	// 12. Search Number
	// Have a target number. Loop through numbers and stop when the target is found.
	var target int

	fmt.Print("Enter a target number: ")
	fmt.Scan(&target)

	for k := 1; k <= 10; k++ {
		if k == target {
			break
		}
		fmt.Println(k)
	}
	// 13. Find First Multiple of 7
	// Loop through numbers and stop when you find the first number divisible by 7.
	for l := 1; l <= 20; l++ {
		if l%7 == 0 {
			fmt.Println("First multiple of 7:", l)
			break
		}
	}

	// 14. Find First Number Greater Than 50
	// Check numbers and stop when one greater than 50 is found.
	for m := 1; m <= 60; m++ {
		if m > 50 {
			fmt.Println("First number greater than 50 found:", m)
			break
		}
	}

	// 15. Sum Until 0
	// Take numbers from the user and keep adding them. Stop when the user enters 0.
	sum99 := 0

	for {
		var numbers int

		fmt.Print("Enter a Number: ")
		fmt.Scan(&numbers)

		if numbers == 0 {
			break
		}
		sum99 = sum99 + numbers
	}
	fmt.Println("Sum:", sum99)

	// 16. Input Until Negative
	// Keep accepting numbers and stop when the user enters a negative number.
	for {
		var untilNegative int

		fmt.Print("Enter a Number: ")
		fmt.Scan(&untilNegative)

		if untilNegative < 0 {
			break
		}
	}
	fmt.Println("Negative Number")

	// 17. Guessing Game
	// Have a predefined number. Keep asking for guesses and stop when the correct number is entered.
	predefinedNumber := 7
	for {
		var guess int

		fmt.Print("Enter a Number: ")
		fmt.Scan(&guess)

		if guess == predefinedNumber {
			break
		}
	}
	fmt.Println("You enter a correct number.")

	// 18. Maximum Attempts
	// Allow a user a maximum of 5 attempts. Use break when the correct answer is found.
	correctNumber := 7
	maxAttempts := 5

	for attemps := 1; attemps <= maxAttempts; attemps++ {
		var guess int

		fmt.Print("Enter a number: ")
		fmt.Scan(&guess)

		if guess == correctNumber {
			fmt.Println("Correct!")
			break
		}

		if attemps == maxAttempts {
			fmt.Println("No more attempts left. The number was", correctNumber)
		}
	}

	// 19. Find First Even Number
	// Search through numbers and stop at the first even number.
	for x := 1; ; x++ {
		if x%2 == 0 {
			fmt.Println("The first even number is: ", x)
			break
		}
	}

	// 20. Find First Odd Number
	// Search through numbers and stop at the first odd number.
	for z := 1; ; z++ {
		if z%2 != 0 {
			fmt.Println("The first odd number is: ", z)
			break
		}
	}
}
