package main

import "fmt"

func main() {
	// For these, practice using fmt.Scan().
	// 1. Ask the user for their name and print a greeting.
	var name string

	fmt.Print("Enter Your Name: ")
	fmt.Scan(&name)

	fmt.Println("Hello", name)

	// 2. Ask the user for their age and print it.
	var age int

	fmt.Print("Enter Your Age: ")
	fmt.Scan(&age)

	fmt.Println(age)

	// 3. Ask the user for their height and print it.
	var height float64

	fmt.Print("Enter Your Height: ")
	fmt.Scan(&height)

	fmt.Println(height)

	// 4. Ask the user for two integers and print their sum.
	var a int
	var b int

	fmt.Print("Enter a num1: ")
	fmt.Scan(&a)
	fmt.Print("Enter a num2: ")
	fmt.Scan(&b)

	fmt.Println(a + b)

	// 5. Ask the user for two integers and print:
	// . Sum
	var c int
	var d int

	fmt.Print("Enter a Num3: ")
	fmt.Scan(&c)
	fmt.Print("Enter a Num4: ")
	fmt.Scan(&d)

	fmt.Println(c + d)

	// . Difference
	var e int
	var f int

	fmt.Print("Enter a Num5: ")
	fmt.Scan(&e)
	fmt.Print("Enter a Num6: ")
	fmt.Scan(&f)

	fmt.Println(e - f)

	// . Product
	var g int
	var h int

	fmt.Print("Enter a Num7: ")
	fmt.Scan(&g)
	fmt.Print("Enter a Num8: ")
	fmt.Scan(&h)

	fmt.Println(g * h)

	// . Quotient
	var i int
	var j int

	fmt.Print("Enter a Num9: ")
	fmt.Scan(&i)
	fmt.Print("Enter a Num10: ")
	fmt.Scan(&j)

	fmt.Println(i / j)

	// . Remainder
	var k int
	var l int

	fmt.Print("Enter a Num11: ")
	fmt.Scan(&k)
	fmt.Print("Ente a Numb12: ")
	fmt.Scan(&l)

	fmt.Println(k % l)
}
