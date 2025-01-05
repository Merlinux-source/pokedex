package main

import (
	"boot.dev-Pokedex/internal/pokeapi"
	"encoding/json"
	"fmt"
)

func commandExplore(conf *CommandConf) error {
	if len(conf.Argv) < 2 || len(conf.Argv) > 3 {
		fmt.Printf("Usage: %s [Location name]\n->\tTo get location names use the Map and Mapb commands to paginate through the location names.\r\n", conf.Argv[0])
		return fmt.Errorf("Invalid command usage")

	}
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
