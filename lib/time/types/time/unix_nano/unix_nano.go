package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	x := t.UnixNano()
	fmt.Println("x:", x)
}
