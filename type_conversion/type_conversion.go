// coversion string into int
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("please rate our pizza between 1 to 5")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Printf("thanks for your %T", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) //remove the error

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Add 1 to your rating ", numRating+1)
	}
}
