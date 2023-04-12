package main

import "fmt"


func commandViewPets(cfg *config, args ...string) error {

	fmt.Println()
	fmt.Println("Your pokemons: ")

	for k, v := range cfg.caughtPokemon {
		fmt.Printf("Pokemon's name: %s\n", k)
		
		fmt.Printf("Height: %d\n", v.Height)
		fmt.Printf("Weight: %d\n", v.Weight)

		fmt.Println("Abilities: ")
		for i := range v.Abilities {
			fmt.Printf("%d. %s\n", i, v.Abilities[i].Ability.Name)
		}

		fmt.Println()
		fmt.Println()
	}

	return nil
}