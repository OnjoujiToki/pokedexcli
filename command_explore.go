package main

import "fmt"

func exploreMap(defaultConfig *configClient, args ...string) error {
	// parse args
	if len(args) != 1 {
		return fmt.Errorf("explore command requires 1 argument")
	}

	locationArea := args[0]
	data, err := defaultConfig.pokeAPIClient.GetLocationArea(locationArea)
	if err != nil {
		return err
	}
	for _, area := range data.PokemonEncounters {
		fmt.Printf("%s\n", area.Pokemon.Name)
	}
	return nil
}
