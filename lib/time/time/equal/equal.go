package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().UTC()
	x := t.Local()
	fmt.Printf("t: %s\nx: %s\nisEqual: %t", t, x, t.Equal(x))
}
