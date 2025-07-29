package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Please input the name of a pokemon")
	}
	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	factor := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %v... \n", name)
	if factor > 40 {
		fmt.Printf("%v escaped! \n", name)
		return nil
	}

	fmt.Printf("%v was caught! \n", name)

	cfg.caughtPokemon[name] = pokemon
	return nil
}
