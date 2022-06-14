package pokemon

type greninja struct {
	Pokemon
}

func newGreninja() Pokemon {
	return &greninja{
		&pokemon{
			name:       "Greninja",
			moves:      []string{"Ice Beam", "Gunk Shot", "U-turn", "Spikes"},
			types:      []string{"Water", "Dark"},
			weaknesses: []string{"Electric", "Fairy", "Fighting", "Grass", "Bug"},
		},
	}
}
