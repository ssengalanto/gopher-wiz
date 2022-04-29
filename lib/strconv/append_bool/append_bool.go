package main

import (
	"fmt"
	"strconv"
)

func main() {
	b := make([]byte, 4)
	b = strconv.AppendBool(b, true)
	fmt.Printf("b: %s, type: %[1]T\n", b)
}
