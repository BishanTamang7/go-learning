package main

import "fmt"

func main() {
	// 1. Create a variable name and store your name. Print it.
	var name string
	name = "Bishan"

	fmt.Println(name)

	// 2. Create variables for age, height, and isStudent. Print all of them.
	var age int
	age = 22
	var height float64 = 1.7
	var isStudent bool = true

	fmt.Println(age)
	fmt.Println(height)
	fmt.Println(isStudent)

	// 3. Declare an int variable without assigning a value. Print its default value.
	var age1 int

	fmt.Println(age1)

	// 4. Declare a string, bool, and float64 without assigning values. Print their default values.
	var name1 string
	var price1 float64
	var grade bool

	fmt.Println(name1)
	fmt.Println(price1)
	fmt.Println(grade)

	// 5. Create age2 := 22, then change it to 23 using =.
	age2 := 22
	age2 = 23

	fmt.Println(age2)

	// 6. Create two variables a := 10 and b := 20. Print both.
	a := 10
	b := 20

	fmt.Println(a)
	fmt.Println(b)

	// 7. Create three variables name7, age7, and country using a single var declaration.
	var (
		name7   string
		age7    int
		country string
	)

	fmt.Println(name7)
	fmt.Println(age7)
	fmt.Println(country)

	// 8. Create price := 100 and quantity := 5. Store their multiplication in a new variable called total.
	price := 100
	quantity := 5

	total := price * quantity

	fmt.Println(total)

	// 9. Create firstName := "Bishan" and lastName := "Tamang". Print both.
	firstName := "Bishan"
	lastName := "Tamang"

	fmt.Println(firstName)
	fmt.Println(lastName)

	// 10. Challenge: Swap the values of p := 10 and q := 20 without using a third variable.
	/*
		Before checking answer
		p := 10
		q := 20

		p = q
		q = p

		fmt.Println(p)
		fmt.Println(q)
	*/

	// After checking answer
	p := 10
	q := 20

	p, q = q, p

	fmt.Println(p)
	fmt.Println(q)
}
