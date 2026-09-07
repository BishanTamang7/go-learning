package main

import "fmt"

func main() {
	// 1. Print 1–10. Print numbers from 1 to 10.
	for a := 1; a <= 10; a++ {
		fmt.Println(a)
	}

	// 2. Print 10–1. Print numbers from 10 down to 1.
	for b := 10; b >= 1; b-- {
		fmt.Println(b)
	}

	// 3. Print 1–100. Print numbers from 1 to 100.
	for c := 1; c <= 100; c++ {
		fmt.Println(c)
	}

	// 4. Print Even Numbers. Print even numbers from 1 to 20.
	for d := 1; d <= 20; d++ {
		if d%2 == 0 {
			fmt.Println(d)
		}
	}

	// 5. Print Odd Numbers. Print odd numbers from 1 to 20.
	for e := 1; e <= 20; e++ {
		if e%2 != 0 {
			fmt.Println(e)
		}
	}

	// 6. Multiples of 5. Print multiples of 5 from 5 to 50.
	for f := 5; f <= 50; f++ {
		if f%5 == 0 {
			fmt.Println(f)
		}
	}

	// 7. Count by 2
	// Print: 2 4 6 8 10 12 ... 20
	for g := 2; g <= 20; g += 2 {
		fmt.Println(g)
	}

	// 8. Count by 3.
	// Print multiples of 3 from 3 to 30.
	for h := 3; h <= 30; h += 3 {
		fmt.Println(h)
	}

	// 9. Print a Word 5 Times
	// Print "Hello" five times.
	for i := 1; i <= 5; i++ {
		fmt.Println("Hello")
	}

	// 10. Print Your Name 10 Times
	// Use a loop to print your name 10 times.
	for j := 1; j <= 10; j++ {
		fmt.Println("Bishan Tamang")
	}

	// 11. Print numbers from 1 to 100 that are divisible by 3.
	for k := 1; k <= 100; k++ {
		if k%3 == 0 {
			fmt.Println(k)
		}
	}

	// 12. Find the sum of numbers from 1 to 10.
	sum := 0

	for l := 1; l <= 10; l++ {
		sum = sum + l
	}
	fmt.Println(sum)

	// 13. Find the sum of even numbers from 1 to 100.
	sum1 := 0

	for m := 1; m <= 100; m++ {
		if m%2 == 0 {
			sum1 = sum1 + m
		}
	}
	fmt.Println(sum1)
}
