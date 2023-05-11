package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now() //show the present time
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2004 15:02:02 monday"))

	createdDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))

}
