package pokemon

import "fmt"

type IPokemonFactory interface {
	MakeWaterPokemon(name string) IPokemon
	MakeGhostPokemon(name string) IPokemon
}

func CreatePokemon(pType string) (IPokemonFactory, error) {
	if pType == "water" {
		return &WaterPokemon{}, nil
	}
	if pType == "ghost" {
		return &GhostPokemon{}, nil
	}
	return nil, fmt.Errorf("wrong pokemon type passed")
}
