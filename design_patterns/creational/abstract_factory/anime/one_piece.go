package anime

type OnePiece interface {
	MakeProtagonist() Protagonist
	MakeAntagonist() Antagonist
}

type onePiece struct{}

func (o *onePiece) MakeProtagonist() Protagonist {
	return &luffy{
		&protagonist{
			name:        "Monkey D. Luffy",
			punchline:   "Meat!",
			affiliation: "Straw Hat Pirates",
			attack:      "Gomu Gomu No King Kong Gatling Gun",
		},
	}
}

func (o *onePiece) MakeAntagonist() Antagonist {
	return &kaidou{
		&antagonist{
			name:        "Kaidou",
			punchline:   "WORORORORO!",
			affiliation: "Beast Pirates",
			attack:      "Thunder Bagua!",
		},
	}
}

// concrete protagonist
type luffy struct {
	Protagonist
}

// concrete antagonist
type kaidou struct {
	Antagonist
}
