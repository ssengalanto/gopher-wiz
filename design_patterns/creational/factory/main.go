package main

import (
	"fmt"
	"github.com/ssengalanto/gopher-wiz/design_patterns/creational/factory/pokemon"
	"strings"
)

func main() {
	waterPokemon, _ := pokemon.CreatePokemon("water")
	ghostPokemon, _ := pokemon.CreatePokemon("ghost")

	greninja := waterPokemon.MakeWaterPokemon("Greninja")
	dragapult := ghostPokemon.MakeGhostPokemon("Dragapult")

	printPokemonDetails(greninja)
	printPokemonDetails(dragapult)
}

func printPokemonDetails(m pokemon.IPokemon) {
	fmt.Printf("Name: %s\n", m.Name())
	fmt.Printf("Moves: %v\n", strings.Join(m.Moves(), ", "))
	fmt.Printf("Types: %v\n", strings.Join(m.Types(), ", "))
	fmt.Printf("Weaknesses: %v\n\n", strings.Join(m.Weaknesses(), ", "))
}
