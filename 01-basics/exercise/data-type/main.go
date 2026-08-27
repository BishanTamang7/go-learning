package main

import "fmt"

func main() {
	// 1. Create an int variable containing your age.
	var age int = 22

	fmt.Println(age)

	// 2. Create a float64 variable containing your height.
	var height float64

	height = 1.7

	fmt.Println(height)

	// 3. Create a string variable containing your name.
	var name string = "Bishan"

	fmt.Println(name)

	// 4. Create a bool variable representing whether you are a student.
	var isStudent bool = true

	fmt.Println(isStudent)

	// 5. Create a variable containing the character 'A'. What data type does Go use for this?
	var grade rune = 'A'

	fmt.Println(grade)

	// 6. Create variables for: Name1, Age1, Height1, Weight, Student status.
	var name1 string = "Bishan"
	var age1 int = 22
	var height1 float64 = 1.7
	var weight float64 = 60
	var isStudent1 bool = true

	fmt.Println(name1)
	fmt.Println(age1)
	fmt.Println(height1)
	fmt.Println(weight)
	fmt.Println(isStudent1)

	// 7. What is wrong with this?
	// var age2 int = 22.5
	// Fix it.
	var age2 int = 22

	fmt.Println(age2)

	// 8. What is wrong with this?
	// var name2 string = 22
	// Fix it.
	var name2 string = "Bishan"

	fmt.Println(name2)

	// 9. What is wrong with this?
	// var isStudent2 bool = "true"
	// Fix it.
	var isStudent2 = true

	fmt.Println(isStudent2)

	// 10. Challenge: Create one variable of each basic type you've learned and print all of them.
	var name3 string = "Bishan"
	var age3 int = 22
	var height3 float64 = 1.7
	var isStudnet3 bool = true
	var grade3 rune = 'A'

	fmt.Println(name3)
	fmt.Println(age3)
	fmt.Println(height3)
	fmt.Println(isStudnet3)
	fmt.Println(grade3)
}
