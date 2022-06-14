package anime

type Protagonist interface {
	Name() string
	Punchline() string
	Affiliation() string
	Save() string
}

type protagonist struct {
	name        string
	punchline   string
	affiliation string
	attack      string
}

func (p protagonist) Name() string {
	return p.name
}

func (p protagonist) Punchline() string {
	return p.punchline
}

func (p protagonist) Affiliation() string {
	return p.affiliation
}

func (p protagonist) Save() string {
	return p.attack
}
