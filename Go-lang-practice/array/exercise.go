//declare the variable by shorthand declaration and accessing
// the element of the array using

package main

import "fmt"

func main() {

	//shorthand declarartion of array

	arr := [4]string{"apple", "mango", "grapes"}

	//Accessing the elements of the array Using for loop
	fmt.Println("Element of the array:")

	for i := 0; i < 3; i++ {
		fmt.Println(arr[i])
	}

}
