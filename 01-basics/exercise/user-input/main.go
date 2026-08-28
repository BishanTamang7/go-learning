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

	// 6. Ask the user for their name1 and age1, then print:
	// Hello Bishan, you are 22 years old.

	var name1 string
	var age1 int

	fmt.Print("Enter Your Name1: ")
	fmt.Scan(&name1)
	fmt.Print("Enter Your Age1: ")
	fmt.Scan(&age1)

	fmt.Println("Hello,", name1, "you are", age1, "years old.")

	// 7. Ask the user for the length and width of a rectangle. Calculate the area.
	var length int
	var width int

	fmt.Print("Enter length of rectangle: ")
	fmt.Scan(&length)
	fmt.Print("Enter Width of rectangle: ")
	fmt.Scan(&width)

	fmt.Println(length * width)

	// 8. Ask the user for the price and quantity of a product. Calculate the total price.
	var price int
	var quantity int

	fmt.Print("Enter a price of a product: ")
	fmt.Scan(&price)
	fmt.Print("Enter a quantity of a product: ")
	fmt.Scan(&quantity)

	fmt.Println(price * quantity)

	// 9. Ask the user for three numbers and calculate their average.
	var num22 int
	var num33 int
	var num44 int

	fmt.Print("Ente a Num22: ")
	fmt.Scan(&num22)
	fmt.Print("Enter a Num33: ")
	fmt.Scan(&num33)
	fmt.Print("Enter a Num44: ")
	fmt.Scan(&num44)

	fmt.Println((num22 + num33 + num44) / 2)

	// 10. Challenge: Ask the user for:
	// . Name7
	// . Age7
	// . Height7
	// . Weight7
	// Then print all information in a formatted way.
	var name7 string
	var age7 int
	var height7 float64
	var weight7 float64

	fmt.Print("Enter your name: ")
	fmt.Scan(&name7)
	fmt.Print("Enter your age: ")
	fmt.Scan(&age7)
	fmt.Print("Enter your height: ")
	fmt.Scan(&height7)
	fmt.Print("Enter your weight: ")
	fmt.Scan(&weight7)

	fmt.Println("Hello, My name is", name7, ". I am", age7, "years old.", ".And i am ", height7, "feet", "And my Weight is", weight7)
}
