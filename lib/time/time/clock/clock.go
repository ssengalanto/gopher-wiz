package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	h, m, s := t.Clock()
	fmt.Printf("h: %d, m: %d, s: %d", h, m, s)
}
