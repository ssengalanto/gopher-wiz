package main

import (
	"fmt"
	"strings"
)

func main() {
	x := strings.ReplaceAll("Go Javascript Gopher Javascript", "Javascript", "Gopher")
	fmt.Println("x:", x)
}
