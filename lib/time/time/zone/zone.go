package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	n, o := t.Zone()
	fmt.Printf("name: %s, offset: %d", n, o)
}
