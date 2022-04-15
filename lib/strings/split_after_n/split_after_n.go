package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Uno,Dos,Tres,Cuatro,Cinco,Seis,Siete"

	x := strings.SplitAfterN(s, ",", 3)
	fmt.Printf("x: %#v\n", x)

	y := strings.SplitAfterN(s, ",", 0)
	fmt.Printf("y: %#v", y)
}
