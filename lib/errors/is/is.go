package main

import (
	"errors"
	"fmt"
)

var ErrNoMana = errors.New("not enough mana")

func main() {
	castSpell := func() error {
		return ErrNoMana
	}

	err := castSpell()
	if errors.Is(err, ErrNoMana) {
		fmt.Println("error: gopher wiz has no mana")
	}
}
