package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	x := t.Second()
	fmt.Println("x:", x)
}
