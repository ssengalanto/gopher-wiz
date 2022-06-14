package pokemon

type Factory interface {
	CreatePokemon(name string) (Pokemon, error)
}

type factory struct {
}

func NewFactory() Factory {
	return &factory{}
}

func (f *factory) CreatePokemon(name string) (Pokemon, error) {
	var val Pokemon

	switch name {
	case "dragapult":
		val = newDragapult()
		break

	case "greninja":
		val = newGreninja()
		break
	}

	return val, nil
}
