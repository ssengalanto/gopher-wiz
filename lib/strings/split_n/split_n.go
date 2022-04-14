package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Uno,Dos,Tres,Cuatro,Cinco,Seis,Siete"

	x := strings.SplitN(s, ",", 3)
	fmt.Printf("x: %#v\n", x)

	y := strings.SplitN(s, ",", 0)
	fmt.Printf("y: %#v", y)
}
