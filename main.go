package main

import (
	"github.com/OnjoujiToki/pokedexcli/internal/pokeapi"
	"time"
)

type configClient struct {
	pokeAPIClient        pokeapi.Client
	nextLocationArea     *string
	previousLocationArea *string
	caughtPokemon        map[string]pokeapi.Pokemon
}

func main() {
	defaultConfig := configClient{
		pokeAPIClient: pokeapi.NewClient(time.Minute * 60),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	startRepl(&defaultConfig)

}
