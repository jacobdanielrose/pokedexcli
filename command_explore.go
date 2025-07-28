package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Please input a location!")
	}
	name := args[0]
	locationResp, err := cfg.pokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationResp.Name)
	fmt.Println("Found Pokemon: ")
	for _, enc := range locationResp.PokemonEncounters {
		fmt.Println(enc.Pokemon.Name)
	}

	return nil
}
