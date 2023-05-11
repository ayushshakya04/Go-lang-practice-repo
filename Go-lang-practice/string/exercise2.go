//Write a program that takes a user’s name and age as input and displays a message on the console saying
// “Hello, [name]! You are [age] years old.

package main

import "fmt"

func main() {
	var (
		str string
		age int
	)
	fmt.Println("enter your name")

	fmt.Scanln(&str)

	fmt.Println("enter your age")
	fmt.Scanln(&age)

	fmt.Println("Hello", str, "! you are", age, "year old.")

}
