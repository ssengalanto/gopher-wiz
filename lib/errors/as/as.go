package main

import (
	"errors"
	"fmt"
)

type GopherWizErr struct {
	spell string
}

func (e *GopherWizErr) Error() string {
	return fmt.Sprintf("not enough mana to cast %s spell", e.spell)
}

func main() {
	var gopherWizErr *GopherWizErr

	castSpell := func(s string) error {
		return &GopherWizErr{spell: s}
	}

	err := castSpell("fmt")
	if errors.As(err, &gopherWizErr) {
		fmt.Println("error:", err)
	}
}
