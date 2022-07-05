package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	y, m, d := t.Date()
	fmt.Printf("%s %d, %d", m, d, y)
}
