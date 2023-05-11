package main

import (
	"fmt"
)

// func main() {
// 	m := map[string]int{ // var map[key]value
// 		"james":           32,
// 		"miss moneypenny": 27,
// 	}
// 	fmt.Println(m)             //print the map
// 	fmt.Println(m["james"])    //print the details of the map values
// 	fmt.Println(m["Barnabas"]) //print the value of that is not in declare in the map
// 	// return the 0 in the output.

// 	v, ok := m["Barnabas"]
// 	fmt.Println(v)
// 	fmt.Println(ok)

// 	if v, ok := m["miss moneypenny"]; ok {
// 		fmt.Println("THis is the if Print", v)
// 	}
// }

// another example

func main() {
	score := make(map[string]int)
	score["bimal"] = 89
	fmt.Println(score)

	score["ron"] = 45
	score["sam"] = 23
	score["vicky"] = 55
	fmt.Println(score)

	fmt.Println(score["sam"]) // do not assign to other variable

	getRonScore := score["ron"] //assign other variable then return the value
	fmt.Println(getRonScore)

	//DELETE Keyword
	delete(score, "vicky")
	fmt.Println(score)

	//FOR LOOP
	//In for loop we use key and value
	for k, v := range score {
		fmt.Printf("score of %v is %v ", k, v)
	}
}
