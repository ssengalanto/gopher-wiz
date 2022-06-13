package main

import (
	"fmt"
	"github.com/ssengalanto/gopher-wiz/design_patterns/creational/factory/pokemon"
	"strings"
)

func main() {
	factory := pokemon.NewPokemonFactory()
	factory.RegisterFactory("dragapult", pokemon.NewDragapult())
	factory.RegisterFactory("greninja", pokemon.NewGreninja())

	dragapult, err := factory.CreatePokemon("dragapult")
	if err != nil {
		panic(err)
	}
	printPokemonDetails(dragapult)

	greninja, err := factory.CreatePokemon("greninja")
	if err != nil {
		return
	}
	printPokemonDetails(greninja)
}

func printPokemonDetails(m pokemon.Pokemon) {
	fmt.Printf("Name: %s\n", m.Name())
	fmt.Printf("Moves: %v\n", strings.Join(m.Moves(), ", "))
	fmt.Printf("Types: %v\n", strings.Join(m.Types(), ", "))
	fmt.Printf("Weaknesses: %v\n\n", strings.Join(m.Weaknesses(), ", "))
}
