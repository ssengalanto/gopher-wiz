package main

import (
	"fmt"
	"time"
)

func main() {
	x, err := time.LoadLocation("Asia/Manila")
	if err != nil {
		panic(err)
	}

	fmt.Println("x:", x)
}
