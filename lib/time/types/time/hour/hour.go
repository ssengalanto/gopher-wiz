package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	x := t.Hour()
	fmt.Println("x:", x)
}
