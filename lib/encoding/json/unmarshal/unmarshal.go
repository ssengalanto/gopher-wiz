package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var data = []byte(`[
	{"Name": "Trafalgar D. Water Law", "Role": ["Former Warlord", "Doctor"]},
	{"Name": "Monkey D. Luffy", "Role": ["Captain", "Rubber man"]}
]`)

	type OnePieceCharacter struct {
		Name string
		Role []string
	}

	var character []OnePieceCharacter

	err := json.Unmarshal(data, &character)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", character)
}
