package pokemon

type Dragapult interface {
	Pokemon
}

type dragapult struct {
	pokemon
}

func NewDragapult() Dragapult {
	return &dragapult{
		pokemon{
			name:       "Dragapult",
			moves:      []string{"Phantom Force", "Protect", "Dragon Darts", "Fly"},
			types:      []string{"Ghost", "Dragon"},
			weaknesses: []string{"Ghost", "Dark", "Dragon", "Fairy", "Ice"},
		},
	}
}
