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

	time := time.Date(1992, 3, 8, 0, 0, 0, 0, time.UTC)

	x := time.In(loc)
	fmt.Println("x:", x)
}
