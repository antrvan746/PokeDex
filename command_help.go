package main

import (
	"fmt"
)

func commandHelp() error {
	fmt.Println()
	fmt.Println()
	fmt.Println("Welcome to PokeDex")
	fmt.Println("Usage: ")
	fmt.Println()
	fmt.Println("hello: Displays a help message")
	fmt.Println("exit: Exit the PokeDex")
	fmt.Println()
	fmt.Println()
	
	return nil
}