package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	x := t.GoString()
	fmt.Println("x:", x)
}
