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
		e := fmt.Errorf("error: %w", err)
		fmt.Println(errors.Unwrap(e))
	}
}
