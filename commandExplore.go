package main

import (
	"boot.dev-Pokedex/internal/pokeapi"
	"encoding/json"
	"fmt"
)

func commandExplore(conf *CommandConf) error {
	exploreLocation := conf.Argv[1]
	if exploreLocation == "" {
		fmt.Printf("Usage: %s [ExploreLocation]\n]", conf.Argv[0])
	}
	resText, err := conf.HttpCache.CacheGet(pokeapi.APIV2_LocationAreaBaseURL + exploreLocation)
	if err != nil {
		fmt.Printf("A error occurred during the exploration: %v\n", err)
		return err
	}
	locationArea := pokeapi.LocationArea{}
	json.Unmarshal(resText, &locationArea)
	fmt.Printf("Pokemon available at %s\r\n", exploreLocation)
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil

}
