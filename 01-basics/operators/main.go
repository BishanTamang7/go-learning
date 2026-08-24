package main

import "fmt"

func main() {
	// Operators are used to perform operations on variables and values.
	// Go supports the following types of operators:
	// Arithmetic Operators. Arithmetic operators are used to perform mathematical operations such as addition, subtraction, multiplication, division, and modulus.
	// + Addition. Adds together two values. For example, 5 + 3 = 8.
	var a int = 10
	var b int = 7

	fmt.Println(a + b) // Output: 17

	// -  Subtraction. Subtracts one value from another. For example, 10 - 5 = 5.
	x := 15
	y := 5

	fmt.Println(x - y) // Output: 10

	// *  Multiplication. Multiplies two values together. For example, 4 * 3 = 12.
	var p int = 10
	var q int = 2

	fmt.Println(p * q) // Output: 20

	// /  Division. Divides one value by another. For example, 10 / 2 = 5.
	m := 20
	n := 4

	fmt.Println(m / n) // Output: 5

	// %  Modulus. Returns the remainder of a division operation. For example, 10 % 3 = 1.
	var r int = 10
	var s int = 3

	fmt.Println(r % s) // Output: 1

	// Relational Operators. Relational operators are used to compare two values and return a boolean result (true or false).
	// ==  Equal to. Checks if two values are equal. For example, 5 == 5 returns true.
	k := 10
	l := 10

	fmt.Println(k == l) // Output: true

	fd := 5
	jb := 10

	fmt.Println(fd == jb) // Output: false

	// !=  Not equal to. Checks if two values are not equal. For example, 5 != 10 returns true.
	var c int = 5
	var d int = 10

	fmt.Println(c != d) // Output: true

	var e int = 7
	var f int = 7

	fmt.Println(e != f) // Output: false

	// >   Greater than. Checks if one value is greater than another. For example, 10 > 5 returns true.
	var g int = 10
	var h int = 5

	fmt.Println(g > h) // Output: true

	var o int = 3
	var rs int = 8

	fmt.Println(o > rs) // Output: false

	// <   Less than. Checks if one value is less than another. For example, 5 < 10 returns true.
	var sw int = 5
	var t int = 10

	fmt.Println(sw < t) // Output: true

	var u int = 15
	var v int = 7

	fmt.Println(u < v) // Output: false

	// >=  Greater than or equal to. Checks if one value is greater than or equal to another. For example, 10 >= 5 returns true.
	var w int = 10
	var x1 int = 10

	fmt.Println(w >= x1) // Output: true

	var y1 int = 5
	var z int = 10

	fmt.Println(y1 >= z) // Output: false

	// <=  Less than or equal to. Checks if one value is less than or equal to another. For example, 5 <= 10 returns true.
	var a1 int = 5
	var b1 int = 10

	fmt.Println(a1 <= b1) // Output: true

	var c1 int = 15
	var d1 int = 7

	fmt.Println(c1 <= d1) // Output: false

	// Logical Operators. Logical operators are used to combine multiple boolean expressions and return a boolean result.
	// &&  Logical AND. Returns true if both expressions are true. For example, (5 > 3) && (10 < 15) returns true.
	var e1 int = 5
	var f1 int = 10

	fmt.Println((e1 > 3) && (f1 < 15)) // Output: true

	var g1 int = 20
	var h1 int = 15

	fmt.Println((g1 > 25) && (h1 < 10)) // Output: false

	// ||  Logical OR. Returns true if at least one of the expressions is true. For example, (5 > 3) || (10 < 5) returns true.
	var i1 int = 5
	var j1 int = 10

	fmt.Println((i1 > 3) || (j1 < 5)) // Output: true

	var k1 int = 20
	var l1 int = 15

	fmt.Println((k1 > 25) || (l1 < 10)) // Output: false

	// !   Logical NOT. Returns true if the expression is false, and false if the expression is true. For example, !(5 > 3) returns false.
	var m1 int = 5
	var n1 int = 10

	fmt.Println(!(m1 > 3)) // Output: false
	fmt.Println(!(n1 < 5)) // Output: true

	// Assignment Operators. Assignment operators are used to assign values to variables. They can also be combined with arithmetic operators to perform operations and assign the result to a variable.
	// =   Assign value. For example, x = 5 assigns the value 5 to the variable x.
	var o1 int = 5
	fmt.Println(o1) // Output: 5

	// +=  Add and assign value. For example, x += 5 is equivalent to x = x + 5.
	var p1 int = 10
	p1 += 5         // Equivalent to p1 = p1 + 5
	fmt.Println(p1) // Output: 15

	// -=  Subtract and assign value. For example, x -= 5 is equivalent to x = x - 5.
	var q1 int = 20
	q1 -= 5         // Equivalent to q1 = q1 - 5
	fmt.Println(q1) // Output: 15

	// *=  Multiply and assign value. For example, x *= 5 is equivalent to x = x * 5.
	var r1 int = 4
	r1 *= 5         // Equivalent to r1 = r1 * 5
	fmt.Println(r1) // Output: 20

	// /=  Divide and assign value. For example, x /= 5 is equivalent to x = x / 5.
	var s1 int = 20
	s1 /= 5         // Equivalent to s1 = s1 / 5
	fmt.Println(s1) // Output: 4

	// %=  Modulus and assign value. For example, x %= 5 is equivalent to x = x % 5.
	var t1 int = 10
	t1 %= 3         // Equivalent to t1 = t1 % 3
	fmt.Println(t1) // Output: 1
}
