package anime

type Antagonist interface {
	Name() string
	Punchline() string
	Affiliation() string
	Destroy() string
}

type antagonist struct {
	name        string
	punchline   string
	affiliation string
	attack      string
}

func (a antagonist) Name() string {
	return a.name
}

func (a antagonist) Punchline() string {
	return a.punchline
}

func (a antagonist) Affiliation() string {
	return a.affiliation
}

func (a antagonist) Destroy() string {
	return a.attack
}
