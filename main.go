package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": cliCommand{
			name:        "help",
			description: "Show help",
			callback: func() error {
				fmt.Println("Help")
				return nil
			},
		},
		"version": cliCommand{
			name:        "version",
			description: "Show version",
			callback: func() error {
				fmt.Println("Version")
				return nil
			},
		},
		"exit": cliCommand{
			name:        "exit",
			description: "Exit",
			callback: func() error {
				fmt.Println("Exit")
				return nil
			},
		},
	}

}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	scanner.Scan()
	for scanner.Text() != "exit" {
		if command, ok := commands[scanner.Text()]; ok {
			command.callback()
		} else {
			fmt.Println("Command not found")
		}
		scanner.Scan()
	}

}
