package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	x := t.Minute()
	fmt.Println("x:", x)
}
