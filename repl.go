package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/antrvan746/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

type cliCommand struct {
	name string
	description string
	callback func(cfg *config) error
}

func getCommands() map[string] cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Exit the PokeDex",
			callback: commandExit,
		},
		"map": {
			name: "map",
			description: "Get the next page of locations",
			callback: commandMapForward,
		},
		"mapb": {
			name: "mapb",
			description: "Get the previous page of locations",
			callback: commandMapBack,
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
		
		command, ok := getCommands()[words[0]]; 
		if ok {
			err := command.callback(cfg)
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

