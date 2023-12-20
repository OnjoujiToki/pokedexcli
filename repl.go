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
	callback    func(defaultConfig *configClient) error
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
			callback: func(defaultConfig *configClient) error {
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
	}
}

func startRepl(defaultConfig *configClient) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands(defaultConfig)
	fmt.Println("Welcome to CLI")

	for {
		fmt.Print(">> ")
		scanner.Scan()
		text := cleanInput(scanner.Text())
		if len(text) == 0 {
			continue
		}
		if command, ok := commands[text[0]]; ok {
			if err := command.callback(defaultConfig); err != nil {
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
