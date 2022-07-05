package main

import (
	"fmt"
	"time"
)

func main() {
	x := time.Now().UTC()
	fmt.Printf("x: %s\n isDST: %t\n", x, x.IsDST())
}
