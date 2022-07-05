package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	x := t.Nanosecond()
	fmt.Println("x:", x)
}
