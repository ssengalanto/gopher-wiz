package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	x := t.UTC()
	fmt.Println("x:", x)
}
