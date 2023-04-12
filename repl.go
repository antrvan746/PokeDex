package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/antrvan746/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	caughtPokemon    map[string]pokeapi.Pokemon
	nextLocationsURL *string
	prevLocationsURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the PokeDex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapForward,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
		"pets": {
			name:        "pets",
			description: "View caught pokemons of pokedex user",
			callback:    commandViewPets,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "Inspect a pokemon",
			callback:    commandInspect,
		},
	}
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		words := cleanInput(input)

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, ok := getCommands()[commandName]
		if ok {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Invalid command")
		}
	}
}

func cleanInput(input string) []string {
	input = strings.ToLower(input)

	words := strings.Fields(input)
	return words
}
