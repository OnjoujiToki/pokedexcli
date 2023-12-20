package main

import "fmt"

func printHelp(defaultConfig *configClient) error {
	fmt.Println("Welcome to the Pokedex CLI helper")
	avaliableCommands := getCommands(defaultConfig)
	for _, command := range avaliableCommands {
		fmt.Printf("%s - %s\n", command.name, command.description)
	}
	// insert an extra line
	fmt.Println()
	return nil
}
