package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Please input the name of a pokemon")
	}
	name := args[0]
	pokemon, ok := cfg.caughtPokemon[name]
	if !ok {
		return errors.New("This pokemon hasn't been caught.")
	}

	fmt.Printf("Name: %v \n", pokemon.Name)
	fmt.Printf("Weight: %v \n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typ := range pokemon.Types {
		fmt.Printf("  - %v \n", typ.Type.Name)
	}
	return nil
}
