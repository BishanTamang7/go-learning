/*
Every Go program starts with a package declaration. The package name is the first line of every Go source file.
The main package is a special package that tells the Go compiler that this is an executable program, rather than a library.
*/
package main

/*
The import statement is used to include other packages in your program.
In this case, we are importing the "fmt" package, which provides functions for formatted I/O, such as printing to the console.
*/
import "fmt"

/*
The main function is the entry point of the program. When you run a Go program, the execution starts from the main function.
In this case, we are using the fmt.Println function to print "Hello, World!" to the console.
*/
func main() {
	/*
		The fmt.Println function prints the specified string to the console, followed by a newline character.
		This is a common way to display output in Go programs.
	*/
	fmt.Println("Hello, World!")
}
