package main

import (
	"fmt"
	"math/rand"
)

func catchPokemon(defaultConfig *configClient, args ...string) error {
	// parse args
	if len(args) != 1 {
		return fmt.Errorf("catch command requires 1 argument")
	}

	pokemon := args[0]
	data, err := defaultConfig.pokeAPIClient.GetPokemon(pokemon)
	if err != nil {
		return err
	}
	experience := data.BaseExperience
	randomNumber := rand.Intn(experience)
	const threshold = 40

	if randomNumber < threshold {
		fmt.Println("Pokemon was not caught")
		return nil
	}
	fmt.Println("Pokemon was caught, its name is", data.Name)
	defaultConfig.caughtPokemon[data.Name] = data

	return nil
}
