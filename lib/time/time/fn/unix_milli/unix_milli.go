package main

import (
	"fmt"
	"time"
)

func main() {
	x := time.Now()
	fmt.Println("x:", x.UnixMilli())
}
