package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a Pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	for _, xp := range pokemon.Base_Experience {
		randumNum := rand.Intn(256)
		if randumNum > xp {
			fmt.Printf("%s was caught!\n", name)
			//TODO: add pokemon to Pokedex
		} else {
			fmt.Printf("%s escaped!\n", name)
		}
	}
	return nil
}
