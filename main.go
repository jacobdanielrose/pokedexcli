package main

import (
	"time"

	"pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	param := ""
	cfg := &config{
		pokeapiClient: pokeClient,
		parameter:     &param,
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	startRepl(cfg)
}
