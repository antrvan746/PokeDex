package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {


	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		
		command, ok := cliCommandMap[input]; 
		if ok {
			command.callback()
		} else {
			fmt.Println("Invalid command")
		}
	}
}