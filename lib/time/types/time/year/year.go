package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	x := t.Year()
	fmt.Println("x:", x)
}
