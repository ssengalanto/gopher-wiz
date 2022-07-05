package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().UTC()
	fmt.Println("t:", t)

	x := t.In(time.Local)
	fmt.Println("x:", x)
}
