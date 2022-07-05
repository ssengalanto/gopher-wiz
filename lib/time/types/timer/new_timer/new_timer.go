package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(time.Second * 3)
	<-timer.C

	fmt.Println("Timer expired.")
}
