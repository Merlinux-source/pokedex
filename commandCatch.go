package main

import (
	"boot.dev-Pokedex/internal/pokeapi"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func commandCatch(conf *CommandConf) error {
	if len(conf.Argv) < 2 || len(conf.Argv) > 3 {
		fmt.Printf("Usage: %s [PokemonToCatch]\n]", os.Args[0])
	}
	pokemonToCatch := conf.Argv[1]
	res, err := conf.HttpCache.CacheGet(pokeapi.APIV2_PokemonBaseURL + pokemonToCatch)
	if err != nil {
		if strings.Contains(err.Error(), "Not Found") {
			fmt.Printf("Pokemon %s not found\r\n", pokemonToCatch)
			return nil
		}
		fmt.Printf("Error occured when attempting to catch pokemon %s\r\nError: %v\r\n", pokemonToCatch, err)
		return nil
	}

	pokemon := pokeapi.Pokemon{}
	fmt.Printf("Throwing a Pokeball at %s...\r\n", pokemonToCatch)
	json.Unmarshal(res, &pokemon)
	chance := 100 - pokemon.BaseExperience
	if chance <= 0 {
		chance = 1
	}
	pokemonIsCaught := rand.Intn(101) > chance
	if pokemonIsCaught {
		fmt.Printf("%s was caught!\r\n", pokemon.Name)
		conf.Pokedex.Add(pokemon)
		return nil
	}
	fmt.Printf("%s escaped!\r\n", pokemon.Name)
	return nil
}
