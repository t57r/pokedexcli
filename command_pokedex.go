package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("There are no pokemons caught yet")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for pokemon := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon)
	}
	return nil
}
