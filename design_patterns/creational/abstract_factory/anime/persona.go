package anime

type Persona interface {
	MakeProtagonist() Protagonist
	MakeAntagonist() Antagonist
}

type persona struct{}

func (p *persona) MakeProtagonist() Protagonist {
	return &joker{
		&protagonist{
			name:        "Ren Amamiya (Joker)",
			punchline:   "Go Phantom Thieves!",
			affiliation: "Phantom Thieves",
			attack:      "Myriad of Truths",
		},
	}
}

func (p *persona) MakeAntagonist() Antagonist {
	return &maruki{
		&antagonist{
			name:        "Takuto Maruki",
			punchline:   "I won't fail...I CAN'T fail!",
			affiliation: "Shujin Academy Councilor",
			attack:      "Full Force!",
		},
	}
}

// concrete protagonist
type joker struct {
	Protagonist
}

// concrete antagonist
type maruki struct {
	Antagonist
}
