package main

import (
	"fmt"
	"log"
)

func printMap(defaultConfig *configClient, args ...string) error {
	// the following codes will show the complete maps name for pokemon world
	response, err := defaultConfig.pokeAPIClient.ListLocationAreas(defaultConfig.nextLocationArea)
	if err != nil {
		log.Fatalln(err)
	}
	defaultConfig.nextLocationArea = response.Next
	defaultConfig.previousLocationArea = response.Previous
	fmt.Println("Location Areas:")
	for _, locationArea := range response.Results {
		fmt.Printf("%s\n", locationArea.Name)
	}

	return nil
}
