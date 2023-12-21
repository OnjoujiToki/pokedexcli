package main

import (
	"fmt"
)

func printMapBack(defaultConfig *configClient, args ...string) error {
	if defaultConfig.previousLocationArea == nil {
		fmt.Println("No more previous location areas")

		return nil
	}
	response, err := defaultConfig.pokeAPIClient.ListLocationAreas(defaultConfig.previousLocationArea)
	if err != nil {
		return err
	}
	defaultConfig.nextLocationArea = response.Next
	defaultConfig.previousLocationArea = response.Previous
	fmt.Println("Previous Location Areas:")
	for _, locationArea := range response.Results {
		fmt.Printf("%s\n", locationArea.Name)
	}
	return nil
}
