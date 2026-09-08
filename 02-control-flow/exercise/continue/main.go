package main

import "fmt"

func main() {
	// 1. Skip 5
	// Print 1–10 but skip 5.
	for a := 1; a <= 10; a++ {
		if a == 5 {
			continue
		}
		fmt.Println(a)
	}

	// 2. Skip 10
	// Print 1–20 but skip 10.
	for b := 1; b <= 20; b++ {
		if b == 10 {
			continue
		}
		fmt.Println(b)
	}

	// 3. Skip Even Numbers
	// Print 1–20 but skip all even numbers.
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// 4. Skip Odd Numbers
	// Print 1–20 but skip all odd numbers.
	for c := 1; c <= 20; c++ {
		if c%2 != 0 {
			continue
		}
		fmt.Println(c)
	}

	// 5. Skip Multiples of 5
	// Print 1–50 but skip numbers divisible by 5.
	for d := 1; d <= 50; d++ {
		if d%5 == 0 {
			continue
		}
		fmt.Println(d)
	}

	// 6. Skip Multiples of 3
	// Print 1–30 but skip numbers divisible by 3.
	for e := 1; e <= 30; e++ {
		if e%3 == 0 {
			continue
		}
		fmt.Println(e)
	}

	// 7. Skip 0
	// Loop from 0 to 10 and don't print 0.
	for g := 0; g <= 10; g++ {
		if g == 0 {
			continue
		}
		fmt.Println(g)
	}

	// 8. Skip Numbers Below 5
	// Loop from 1 to 10 and skip numbers below 5.
	for j := 1; j <= 10; j++ {
		if j < 5 {
			continue
		}
		fmt.Println(j)
	}

	// 9. Skip Numbers Above 5
	// Loop from 1 to 10 and skip numbers above 5.
	for k := 1; k <= 10; k++ {
		if k > 5 {
			continue
		}
		fmt.Println(k)
	}

	// 10. Skip 3 and 7
	// Print 1–10 but don't print 3 or 7.
	for l := 1; l <= 10; l++ {
		if l == 3 || l == 7 {
			continue
		}
		fmt.Println(l)
	}

	// 11. Only Even Numbers
	// Loop from 1 to 100 and use continue to skip odd numbers.
	for m := 1; m <= 100; m++ {
		if m%2 != 0 {
			continue
		}
		fmt.Println(m)
	}

	// 12. Only Odd Numbers
	// Loop from 1 to 100 and use continue to skip even numbers.
	for o := 1; o <= 100; o++ {
		if o%2 == 0 {
			continue
		}
		fmt.Println(o)
	}

	// 13. Skip Multiples of 10
	// Print 1–100 while skipping multiples of 10.
	for n := 1; n <= 100; n++ {
		if n%10 == 0 {
			continue
		}
		fmt.Println(n)
	}

	// 14. Skip Numbers Divisible by 3
	// Print 1–50 but skip numbers divisible by 3.
	for p := 1; p <= 50; p++ {
		if p%3 == 0 {
			continue
		}
		fmt.Println(p)
	}

	// 15. Skip Numbers Divisible by 3 and 5
	// Print 1–100 but skip numbers divisible by both 3 and 5.
	for q := 1; q <= 100; q++ {
		if q%3 == 0 && q%5 == 0 {
			continue
		}
		fmt.Println(q)
	}

	// 16. Skip Negative Numbers
	// Take several numbers from the user. Use continue to skip negative numbers.
	for r := 1; r <= 10; r++ {
		var r int

		fmt.Print("Enter a number: ")
		fmt.Scan(&r)

		if r < 0 {
			continue
		}
		fmt.Println(r)
	}

	// 17. Skip Zero
	// Take numbers from the user and skip zero values.
	for u := 1; u <= 10; u++ {
		var numbers7 int

		fmt.Print("Enter a number: ")
		fmt.Scan(&numbers7)

		if numbers7 == 0 {
			continue
		}
		fmt.Println(numbers7)
	}

	// 18. Skip Invalid Scores
	// Take several scores. Skip scores outside the range 0–100.
	for z := 1; z <= 100; z++ {
		var score int

		fmt.Print("Enter a Score: ")
		fmt.Scan(&score)

		if z < 0 || z > 100 {
			continue
		}
		fmt.Println(score)
	}

	// 19. Skip Small Numbers
	// Print numbers from 1–50 but skip numbers less than 10.
	for t := 1; t <= 50; t++ {
		if t < 10 {
			continue
		}
		fmt.Println(t)
	}

	// 20. Skip Large Numbers
	// Print numbers from 1–50 but skip numbers greater than 40.
	for s := 1; s <= 50; s++ {
		if s > 40 {
			continue
		}
		fmt.Println(s)
	}
}
