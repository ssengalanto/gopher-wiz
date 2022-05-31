package main

import (
	"fmt"
	"time"
)

func main() {
	x, err := time.Parse("Jan 2, 2006 at 3:04pm (MST)", "Mar 8, 1992 at 7:54pm (PST)")
	if err != nil {
		panic(err)
	}
	fmt.Println("x:", x)

	y, err := time.Parse("Jan 2, 2006", "Mar 8, 1992")
	if err != nil {
		panic(err)
	}
	fmt.Println("y:", y)

	z, err := time.Parse(time.RFC3339, "1992-03-08T00:00:00Z")
	if err != nil {
		panic(err)
	}
	fmt.Println("z:", z)
}
