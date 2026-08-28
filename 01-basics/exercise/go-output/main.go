package main

import "fmt"

func main() {
	// 1. Print "Hello" using fmt.Print().
	fmt.Print("Hello")

	// 2. Print "Hello" and "World" on the same line using fmt.Print().
	fmt.Print("Hello ")
	fmt.Print("World")

	// 3. Print "Hello" and "World" on separate lines using fmt.Print() and \n.
	fmt.Print("Hello\n")
	fmt.Print("World")

	// 4. Print your name and age using fmt.Println().
	name := "Bishan"
	age := 22

	fmt.Println(name, age)
	// 5. Print three values using one fmt.Println().
	a := 7
	b := 8
	c := 9

	fmt.Println(a, b, c)

	// 6. Use fmt.Printf() to print:
	// My name is Bishan.
	// Use %s.
	var name1 string = "Bishan"

	fmt.Printf("My name is %s.", name1)

	// 7. Use %d to print your age.
	age1 := 22

	fmt.Printf("Age: %d", age1)

	// 8. Use %f to print your height.
	height := 1.7

	fmt.Printf("Height: %f", height)

	// 9. Use %t to print your student status.
	isStudent := true

	fmt.Printf("IsStudent: %t", isStudent)

	// 10. Challenge: Create:
	// name2 := "Abi"
	// age2 := 21
	// height2 := 1.9
	// isStudent2 := true
	// Print everything using one fmt.Printf().
	name2 := "Abi"
	age2 := 21
	height2 := 1.9
	isStudent2 := true

	fmt.Printf("My name is %s. I am %d years old. My Height is %fm. Student: %t.\n", name2, age2, height2, isStudent2)
}
