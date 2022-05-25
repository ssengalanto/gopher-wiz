package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.Tick(1 * time.Second)
	for t := range ticker {
		h, m, s := t.Clock()
		fmt.Printf("hour: %d: minute: %d: second:%d\n", h, m, s)
	}
}
