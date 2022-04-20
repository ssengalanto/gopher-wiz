package main

import (
	"fmt"
	"strings"
)

func main() {
	x := strings.TrimPrefix("-0_Gopher Wiz_0-", "-0_")
	fmt.Println("x:", x)
}
