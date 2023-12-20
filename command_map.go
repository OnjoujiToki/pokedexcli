package main

import (
	"fmt"
	"log"
)

func printMap(defaultConfig *configClient) error {
	// the following codes will show the complete maps name for pokemon world
	pokemonAPIClient := defaultConfig.pokeAPIClient
	response, err := pokemonAPIClient.ListLocationAreas(defaultConfig.nextLocationArea)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Location Areas:")
	for _, locationArea := range response.Results {
		fmt.Printf("%s\n", locationArea.Name)
	}
	defaultConfig.nextLocationArea = response.Next

	return nil
}
