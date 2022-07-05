package main

import (
	"fmt"
	"time"
)

func main() {
	x := time.Now()
	fmt.Println("x:", x.Add(1*time.Hour+10).Format(time.Kitchen))
}
