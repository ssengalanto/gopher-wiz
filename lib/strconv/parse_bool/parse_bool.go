package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := []string{"1", "t", "T", "TRUE", "true", "True", "0", "f", "F", "FALSE", "false", "False", "yes", "no"}

	for _, v := range s {
		x, err := strconv.ParseBool(v)
		if err != nil {
			fmt.Printf("value: %v, error: %v\n", x, err)
			continue
		}
		fmt.Printf("value: %t, type: %[1]T\n", x)
	}
}
