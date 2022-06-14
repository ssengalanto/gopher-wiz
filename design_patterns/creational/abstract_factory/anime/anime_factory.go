package anime

import "fmt"

type AnimeFactory interface {
	MakeProtagonist() Protagonist
	MakeAntagonist() Antagonist
}

type animeFactory struct {
}

func CreateAnime(anime string) (AnimeFactory, error) {
	if anime == "onepiece" {
		return &onePiece{}, nil
	}
	if anime == "persona" {
		return &persona{}, nil
	}

	return nil, fmt.Errorf("invalid payload anime: %s", anime)
}
