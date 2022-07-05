package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	x := t.Month()
	fmt.Println("x:", x)
}
