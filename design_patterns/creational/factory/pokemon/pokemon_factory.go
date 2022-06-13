package pokemon

import "fmt"

type PokemonFactory interface {
	RegisterFactory(name string, pokemon Pokemon)
	CreatePokemon(name string) (Pokemon, error)
}

type pokemonFactory struct {
	factories map[string]Pokemon
}

func NewPokemonFactory() PokemonFactory {
	return &pokemonFactory{make(map[string]Pokemon)}
}

func (p *pokemonFactory) RegisterFactory(name string, pokemon Pokemon) {
	p.factories[name] = pokemon
}

func (p *pokemonFactory) CreatePokemon(name string) (Pokemon, error) {
	val, ok := p.factories[name]

	if !ok {
		return nil, fmt.Errorf("invalid payload name: %s", name)
	}

	return val, nil
}
