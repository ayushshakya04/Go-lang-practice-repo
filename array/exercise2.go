package main

import "fmt"

func main() {

	arr := [5]int{1, 2, 3, 4, 5}
	//size of array
	n := 5
	//declare a variable to store the sum
	sum := 0

	for i := 0; i < n; i++ {
		//adding the values of array to the variable sum
		sum += (arr[i])

	}

	//declare a variable avg to find the average
	avg := (float32(sum)) / (float32(n))
	fmt.Println(avg)

}
