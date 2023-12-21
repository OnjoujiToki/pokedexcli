package main

import (
	"fmt"
)

func checkPokemon(defaultConfig *configClient, args ...string) error {
	// parse args
	if len(args) != 1 {
		return fmt.Errorf("catch command requires 1 argument")
	}

	// check all pokemons the user own
	for pokemon := range defaultConfig.caughtPokemon {
		if pokemon == args[0] {
			fmt.Println("Pokemon was caught")
			return nil
		}
	}
	fmt.Println("Pokemon was not caught")
	return nil
}
