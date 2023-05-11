package main

import (
	"fmt"
)

func main() {

	//Declaring a single variable
	//var name type
	var age int // variable declaration
	fmt.Println("My age is", age)

	//Declaring a variable with an initial value
	//var name type = initialvalue
	var age1 int = 29 //variable declaration with initial value.\
	fmt.Println("my age is", age1)

	//Type inference :
	//if a variable has a inital value, Go will automatically be
	//able to infer the type of that variable using that initial value.Hence if a variable
	//has an initial value, the "type" in the variable declarartion can be removed.

	//var name = initialvalue
	var age2 = 29
	fmt.Println("my age is", age2)

	//multiple variable declarartion
	//var name1,name2type=initalvalue1,initialvalue2

	var width, height int = 100, 50 // declaring multiple variables
	fmt.Println("width is", width, "height is", height)

	var width1, height1 int
	fmt.Println("width is", width, "height is", width1)
	width1 = 100
	height1 = 50
	fmt.Println("new width is", width, "new height is", height1)

	//var(
	//name1=initalvalue1
	//name2=initalvalue2
	//)

	var (
		name    = "naveen"
		age3    = 29
		height3 = int
	)
	fmt.Println("my name is", name, ",age is ", age3, "and height", height3)

	//short hand declaration
	//name:= initialvalue
	count := 10
	fmt.Println("count =", count)

	//declare mutliple variable
	name, age3 = "naveen", 29
	fmt.Println("my name is", name, "age is", age)
}
