package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	y, w := t.ISOWeek()
	fmt.Printf("year: %d, week: %d", y, w)
}
