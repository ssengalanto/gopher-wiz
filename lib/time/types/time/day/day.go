package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	x := t.Day()
	fmt.Println("x:", x)
}
