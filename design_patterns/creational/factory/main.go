package main

import (
	"fmt"
	"github.com/ssengalanto/gopher-wiz/design_patterns/creational/factory/pokemon"
	"strings"
)

func main() {
	pokeFactory := pokemon.NewFactory()

	dragapult, err := pokeFactory.CreatePokemon("dragapult")
	if err != nil {
		panic(err)
	}
	fmt.Println("--- DRAGAPULT ---")
	printPokemonDetails(dragapult)

	greninja, err := pokeFactory.CreatePokemon("greninja")
	if err != nil {
		return
	}
	fmt.Println("--- GRENINJA ---")
	printPokemonDetails(greninja)
}

func printPokemonDetails(m pokemon.Pokemon) {
	fmt.Printf("Name: %s\n", m.Name())
	fmt.Printf("Moves: %v\n", strings.Join(m.Moves(), ", "))
	fmt.Printf("Types: %v\n", strings.Join(m.Types(), ", "))
	fmt.Printf("Weaknesses: %v\n\n", strings.Join(m.Weaknesses(), ", "))
}
