package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	x := t.Weekday()
	fmt.Println("x:", x)
}
