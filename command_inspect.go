package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you need to put a pokemon's name")
	}

	name := args[0]
	pokemon, ok := cfg.caughtPokemon[name]

	if !ok {
		return errors.New("you have not catch this pokemon")
	}

	fmt.Printf("Pokemon's name: %s\n", pokemon.Name)

	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Println("Abilities: ")
	for i := range pokemon.Abilities {
		fmt.Printf("%d. %s\n", i, pokemon.Abilities[i].Ability.Name)
	}

	fmt.Println("Stats: ")
	for i := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", pokemon.Stats[i].Stat.Name, pokemon.Stats[i].BaseStat)
	}

	fmt.Println("Types: ")
	for i := range pokemon.Types {
		fmt.Printf("  %d. %s\n", i, pokemon.Types[i].Type.Name)
	}

	fmt.Println()
	fmt.Println()

	return nil
}
