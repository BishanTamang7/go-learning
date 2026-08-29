package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 1. Convert:
	age := 22
	// from int to float64.
	agefloat := float64(age)

	fmt.Println(reflect.TypeOf(age))
	fmt.Println(reflect.TypeOf(agefloat))

	// 2. Convert:
	height := 5.9
	// from float64 to int.
	heightint := int(height)

	fmt.Println(reflect.TypeOf(height))
	fmt.Println(reflect.TypeOf(heightint))

	// 3. Convert:
	price := 99.99
	// from float64 to int.
	priceint := int(price)

	fmt.Println(reflect.TypeOf(price))
	fmt.Println(reflect.TypeOf(priceint))

	// 4. Convert:
	number := 100
	// from int to float64 and calculate:
	numberfloat := float64(number)

	fmt.Println(reflect.TypeOf(number))
	fmt.Println(reflect.TypeOf(numberfloat))

	// 5. Given:
	a := 10
	b := 5.5
	// Add them together using type conversion.
	add := float64(a) + b

	fmt.Println(add)

	// 6. Given:
	quantity := 5
	price2 := 19.99
	// Calculate the total price using type conversion.
	calculate := float64(quantity) * price2

	fmt.Println(calculate)

	// 7.Given:
	totalMarks := 450
	// Convert it to float64 and calculate an average for 5 subjects.
	float := float64(totalMarks)

	fmt.Println(float / 5)

	// 8. Given:
	temperature := 37.8
	// Convert it to int. Print both values.
	int1 := int(temperature)

	fmt.Println(temperature)
	fmt.Println(int1)

	// 9. Given:
	distance := 100
	time := 2.5
	// Use type conversion where necessary.
	// Calculate speed:
	distanceint := float64(distance)

	speed := distanceint / time

	fmt.Println(speed)

	// 10. Challenge: Given:
	var num1 int = 10
	var num2 float64 = 3.5
	// Calculate:
	// 10 + 3.5
	// 10 - 3.5
	// 10 × 3.5
	// 10 ÷ 3.5
	// using explicit type conversion.

	num3 := float64(num1)

	fmt.Println(num3 + num2)
	fmt.Println(num3 - num2)
	fmt.Println(num3 * num2)
	fmt.Println(num3 / num2)

}
