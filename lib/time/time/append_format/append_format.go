package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	txt := []byte("Kitchen Time: ")

	x := t.AppendFormat(txt, time.Kitchen)
	fmt.Println("x:", string(x))
}
