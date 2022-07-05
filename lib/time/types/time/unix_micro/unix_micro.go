package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	x := t.UnixMicro()
	fmt.Println("x:", x)
}
