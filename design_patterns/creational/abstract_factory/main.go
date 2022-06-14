package main

import (
	"fmt"
	"github.com/ssengalanto/gopher-wiz/design_patterns/creational/abstract_factory/anime"
)

func main() {
	onepiece, _ := anime.CreateAnime("onepiece")
	persona, _ := anime.CreateAnime("persona")

	luffy := onepiece.MakeProtagonist()
	kaidou := onepiece.MakeAntagonist()

	joker := persona.MakeProtagonist()
	maruki := persona.MakeAntagonist()

	fmt.Printf("--- ONE PIECE ---")
	printCharacterDetails(luffy)
	printVillainDetails(kaidou)

	fmt.Printf("\n --- PERSONA 5 ---")
	printCharacterDetails(joker)
	printVillainDetails(maruki)
}

func printCharacterDetails(p anime.Protagonist) {
	fmt.Printf("\nName: %s\n", p.Name())
	fmt.Printf("PunchLine: %s\n", p.Punchline())
	fmt.Printf("Affiliation: %s\n", p.Affiliation())
	fmt.Printf("Save: %s\n", p.Save())
}

func printVillainDetails(a anime.Antagonist) {
	fmt.Printf("\nName: %s\n", a.Name())
	fmt.Printf("PunchLine: %s\n", a.Punchline())
	fmt.Printf("Affiliation: %s\n", a.Affiliation())
	fmt.Printf("Destroy: %s\n", a.Destroy())
}
