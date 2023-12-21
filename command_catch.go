package main

import "fmt"

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
	fmt.Printf("You catched %s\n", data.Name)
	for _, ability := range data.Abilities {
		fmt.Printf("Ability: %s\n", ability.Ability.Name)
	}
	fmt.Printf("Base experience: %d\n", data.BaseExperience)
	fmt.Printf("Height: %d\n", data.Height)
	fmt.Printf("Weight: %d\n", data.Weight)
	fmt.Printf("Order: %d\n", data.Order)
	fmt.Printf("ID: %d\n", data.ID)
	fmt.Printf("Is default: %t\n", data.IsDefault)
	fmt.Printf("Location area encounters: %s\n", data.LocationAreaEncounters)
	fmt.Printf("Name: %s\n", data.Name)
	fmt.Printf("Species: %s\n", data.Species.Name)
	fmt.Printf("Sprites: %s\n", data.Sprites.FrontDefault)

	return nil
}
