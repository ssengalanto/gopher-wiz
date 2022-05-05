package pokemon

type IPokemon interface {
	Name() string
	Moves() []string
	Types() []string
	Weaknesses() []string
}

type Pokemon struct {
	name       string
	moves      []string
	types      []string
	weaknesses []string
}

func (p *Pokemon) Name() string {
	return p.name
}

func (p *Pokemon) Moves() []string {
	return p.moves
}

func (p *Pokemon) Types() []string {
	return p.types
}

func (p *Pokemon) Weaknesses() []string {
	return p.weaknesses
}

type WaterPokemon struct {
	Pokemon
}

func (p *Pokemon) MakeWaterPokemon(name string) IPokemon {
	return &WaterPokemon{
		Pokemon{
			name:       name,
			moves:      []string{"Hydro Pump", "Surf"},
			types:      []string{"Water"},
			weaknesses: []string{"Grass", "Lightning"},
		},
	}
}

type GhostPokemon struct {
	Pokemon
}

func (p *Pokemon) MakeGhostPokemon(name string) IPokemon {
	return &GhostPokemon{
		Pokemon{
			name:       name,
			moves:      []string{"Shadow Ball", "Hex"},
			types:      []string{"Ghost"},
			weaknesses: []string{"Ghost", "Dark"},
		},
	}
}
