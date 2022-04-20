package main

import (
	"fmt"
	"strings"
)

func main() {
	x := strings.TrimPrefix("-0_Gopher Wiz_0-", "_0-")
	fmt.Println("x:", x)
}
