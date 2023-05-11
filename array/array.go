package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to array")
	//Example 1

	//decalre the array
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "banana"
	fruitList[2] = "mango"
	fruitList[3] = "peach"
	fmt.Println(fruitList)
	fmt.Println(len(fruitList)) // find the lenght of array

	//one more way to declare the array
	//Example 2

	var vegList = [4]string{"potato", "tomato", "carrot", "onion"}

	fmt.Println(vegList)
	fmt.Println(len(vegList))

	//example : 3

	// var x [5]int
	// fmt.Println(x)
	// x[3] = 42
	// fmt.Println(x)
	// fmt.Println((len(x)))   //this line say the same thing what first code saying
	// var arr = [...]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 11, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	// fmt.Println(arr)

	//Example 4:
}
