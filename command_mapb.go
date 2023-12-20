package main

import (
	"fmt"
	"log"
)

func printMapBack(defaultConfig *configClient) error {
	fmt.Println("Map back")
	response, err := defaultConfig.pokeAPIClient.ListLocationAreas(defaultConfig.previousLocationArea)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Previous Location Areas:")
	for _, locationArea := range response.Results {
		fmt.Printf("%s\n", locationArea.Name)
	}
	return nil
}
