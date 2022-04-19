package main

import (
	"fmt"
	"strings"
)

func main() {
	x := strings.TrimRight("-_Gopher Wiz_-", "-_")
	fmt.Println("x:", x)
}
