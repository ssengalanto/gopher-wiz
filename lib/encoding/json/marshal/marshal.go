package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type OnePieceCharacter struct {
		ID   int
		Name string
		Role []string
	}

	group := OnePieceCharacter{
		ID:   1,
		Name: "Trafalgar D. Water Law",
		Role: []string{"Former Warlord", "Doctor"},
	}

	x, err := json.Marshal(group)
	if err != nil {
		panic(err)
	}

	fmt.Println("x:", string(x))
}
