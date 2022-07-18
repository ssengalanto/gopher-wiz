package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	x := json.Valid([]byte(`{"name": "Ssen Galanto"}`))
	fmt.Println("x:", x)

	y := json.Valid([]byte(`{"name": Ssen Galanto}`))
	fmt.Println("y:", y)
}
