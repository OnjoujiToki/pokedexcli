package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*configClient, ...string) error
}

func getCommands(defaultConfig *configClient) map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Show help",
			callback:    printHelp,
		},
		"version": {
			name:        "version",
			description: "Show version",
			callback: func(defaultConfig *configClient, args ...string) error {
				fmt.Println("v0.0.1")
				return nil
			},
		},
		"exit": {
			name:        "exit",
			description: "Exit the CLI program",
			callback:    printExit,
		},
		"map": {
			name:        "map",
			description: "Show map",
			callback:    printMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Show map back",
			callback:    printMapBack,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "Explore the map, list the pokemon encounters",
			callback:    exploreMap,
		},
		"catch": {
			name:        "catch {pokemon}",
			description: "Catch a pokemon",
			callback:    catchPokemon,
		},
		"check": {
			name:        "check {pokemon}",
			description: "Check if a pokemon was caught",
			callback:    checkPokemon,
		},
	}
}

func startRepl(defaultConfig *configClient) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to CLI")

	for {

		fmt.Print(">> ")
		scanner.Scan()
		text := cleanInput(scanner.Text())
		if len(text) == 0 {
			continue
		}
		args := []string{}
		if len(text) > 1 {
			args = text[1:]
		}
		commands := getCommands(defaultConfig)
		if command, ok := commands[text[0]]; ok {
			if err := command.callback(defaultConfig, args...); err != nil {
				if err != nil {
					return
				} else {
					fmt.Println("Command not found")
				}
			}
		} else {
			fmt.Println("Command not found")
		}

	}
}

func cleanInput(input string) []string {
	return strings.Fields(strings.ToLower(input))
}
