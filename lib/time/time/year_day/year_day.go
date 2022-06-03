package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	x := t.YearDay()
	fmt.Println("x:", x)
}
