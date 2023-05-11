//BASIC CODE FOR EXPLAINING THE SYNTAX OF THE CODE

// package    	// Package declaration: All Go programs start with a package declaration.
// The package "main" is a special package that defines a standalone executable program.

import "fmt" // Import statement:  The import statement is used to import packages into the program.
// In this case, the "fmt" package is imported, which provides formatting and printing functions.

// Main function:The main function is the entry point of the program, and all Go programs must have a main function.
//The code inside the main function is executed when the program is run.
func main1() {
	fmt.Println("hello world")
	foo()

	fmt.Println("hello")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

}

func foo() {
	panic("hell")
}
