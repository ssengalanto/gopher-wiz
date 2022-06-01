package main

import (
	"fmt"
	"time"
)

func main() {
	d := time.Date(1992, 3, 8, 0, 0, 0, 0, time.UTC)

	x := d.AddDate(45, 7, 13)
	fmt.Println("x:", x)
}
