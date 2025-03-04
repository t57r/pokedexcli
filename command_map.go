package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config) error {
	return ExecuteAndListLocations(cfg, cfg.nextLocationsURL)
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}
	return ExecuteAndListLocations(cfg, cfg.prevLocationsURL)
}

func ExecuteAndListLocations(cfg *config, url *string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(url)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
