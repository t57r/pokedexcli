package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	exploredArea := args[0]
	fmt.Printf("Exploring %s...\n", exploredArea)
	exploredLocationResp, err := cfg.pokeapiClient.ExploreLocation(&exploredArea)
	if err != nil {
		return err
	}

	if len(exploredLocationResp.PokemonEncounters) == 0 {
		fmt.Println("No pokemons found")
		return nil
	}

	fmt.Println("Found Pokemon:")
	for _, enc := range exploredLocationResp.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}
	return nil
}
