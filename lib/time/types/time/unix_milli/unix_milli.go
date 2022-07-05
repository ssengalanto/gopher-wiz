package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	x := t.UnixMilli()
	fmt.Println("x:", x)
}
