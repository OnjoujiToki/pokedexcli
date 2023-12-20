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
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Show help",
			callback: func() error {
				printHelp()
				return nil
			},
		},
		"version": {
			name:        "version",
			description: "Show version",
			callback: func() error {
				fmt.Println("Version")
				return nil
			},
		},
		"exit": {
			name:        "exit",
			description: "Exit the CLI program",
			callback: func() error {
				printExit()
				return nil
			},
		},
	}
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	fmt.Println("Welcome to CLI")

	for {
		fmt.Print(">> ")
		scanner.Scan()
		text := cleanInput(scanner.Text())
		if len(text) == 0 {
			continue
		}
		if command, ok := commands[text[0]]; ok {
			if err := command.callback(); err != nil {
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
