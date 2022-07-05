package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(1992, 3, 8, 12, 15, 30, 987654321, time.UTC)

	trunc := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		time.Minute,
		time.Hour,
	}

	for _, d := range trunc {
		fmt.Printf("(%5s) = %s\n", d, t.Truncate(d).Format("15:04:05.999999999"))
	}
}
