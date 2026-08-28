package main

import "fmt"

func main() {
	// 1. Given:
	a := 20
	b := 5

	// Calculate addition.
	fmt.Println(a + b) // Output: 25
	// 2. Calculate subtraction using the same variables.
	fmt.Println(a - b) // Output: 15

	// 3. Calculate multiplication.
	fmt.Println(a * b) // Output: 100

	// 4. Calculate division.
	fmt.Println(a / b) // Output: 4

	// 5. Calculate the remainder using %.
	fmt.Println(a % b) // Ouptu: 0

	// 6. Given:
	c := 10
	d := 20

	// Check whether c == d.
	// 7. Check whether c != d.
	fmt.Println(c != d) // Ouput: true

	// 8. Check whether c > d, c < d, c >= d, and c <= d.
	fmt.Println(c > d)  // Output: false
	fmt.Println(c < d)  // Output: true
	fmt.Println(c >= d) // Ouput: false
	fmt.Println(c <= d) // Ouput: true

	// 9. Given:
	age := 22
	isStudent := true

	// Use &&, ||, and ! to create three different logical expressions.
	// Using AND Gate (&&)
	expression1 := age >= 18 && isStudent
	fmt.Println(expression1) // Ouput: true

	// Using OR Gate (||)
	expression2 := age >= 18 || isStudent
	fmt.Println(expression2) // Ouput: true

	// Uisng NOT Gate (!)
	expression3 := !isStudent
	fmt.Println(expression3) // Ouput: false

	// 10. Challenge: Given:
	number := 25

	// Use % to determine whether number is divisible by 5.
	fmt.Println(number%5 == 0) // Ouput: ture
}
