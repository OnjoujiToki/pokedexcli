package main

import "github.com/OnjoujiToki/pokedexcli/internal/pokeapi"

type configClient struct {
	pokeAPIClient        pokeapi.Client
	nextLocationArea     *string
	previousLocationArea *string
}

func main() {
	defaultConfig := configClient{
		pokeAPIClient: pokeapi.NewClient(),
	}
	startRepl(&defaultConfig)

}
