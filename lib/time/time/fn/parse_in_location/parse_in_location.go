package main

import (
	"fmt"
	"time"
)

func main() {
	loc, err := time.LoadLocation("Asia/Manila")
	if err != nil {
		panic(err)
	}

	x, err := time.ParseInLocation("Jan 2, 2006 at 3:04pm (MST)", "Mar 8, 1992 at 7:54pm (PST)", loc)
	if err != nil {
		panic(err)
	}
	fmt.Println("x:", x)

	y, err := time.ParseInLocation("Jan 2, 2006", "Mar 8, 1992", loc)
	if err != nil {
		panic(err)
	}
	fmt.Println("y:", y)

	z, err := time.ParseInLocation(time.RFC3339, "1992-03-08T00:00:00Z", loc)
	if err != nil {
		panic(err)
	}
	fmt.Println("z:", z)
}
