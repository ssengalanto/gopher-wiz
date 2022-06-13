package pokemon

type Greninja interface {
	Pokemon
}

type greninja struct {
	pokemon
}

func NewGreninja() Greninja {
	return &greninja{
		pokemon{
			name:       "Greninja",
			moves:      []string{"Ice Beam", "Gunk Shot", "U-turn", "Spikes"},
			types:      []string{"Water", "Dark"},
			weaknesses: []string{"Electric", "Fairy", "Fighting", "Grass", "Bug"},
		},
	}
}
