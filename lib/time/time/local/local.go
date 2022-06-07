package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	x := t.Local()
	fmt.Println("x:", x)
}
