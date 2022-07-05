package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	x := t.Unix()
	fmt.Println("x:", x)
}
