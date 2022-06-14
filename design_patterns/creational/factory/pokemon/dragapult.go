package pokemon

type dragapult struct {
	Pokemon
}

func newDragapult() Pokemon {
	return &dragapult{
		&pokemon{
			name:       "Dragapult",
			moves:      []string{"Phantom Force", "Protect", "Dragon Darts", "Fly"},
			types:      []string{"Ghost", "Dragon"},
			weaknesses: []string{"Ghost", "Dark", "Dragon", "Fairy", "Ice"},
		},
	}
}
