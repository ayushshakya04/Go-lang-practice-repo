// Create a program that takes a number as input and displays the multiplication table of that number up to 10.

package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for i := 0; i <= 10; i++ {
		fmt.Println(n, "x", i, "=", n*i)
	}

}
