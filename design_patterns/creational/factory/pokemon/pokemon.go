package pokemon

type Pokemon interface {
	Name() string
	Moves() []string
	Types() []string
	Weaknesses() []string
}

type pokemon struct {
	name       string
	moves      []string
	types      []string
	weaknesses []string
}

func (p *pokemon) Name() string {
	return p.name
}

func (p *pokemon) Moves() []string {
	return p.moves
}

func (p *pokemon) Types() []string {
	return p.types
}

func (p *pokemon) Weaknesses() []string {
	return p.weaknesses
}
