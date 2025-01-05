package main

import (
	"fmt"
)

func commandInspect(conf *CommandConf) error {
	if len(conf.Argv) < 2 || len(conf.Argv) > 3 {
		fmt.Printf("Usage: %s [pokemon name]\n", conf.Argv[0])
		return fmt.Errorf("Invalid command usage")
	}

	pokemonName := conf.Argv[1]
	pokemon, err := conf.Pokedex.Get(pokemonName)
	if err != nil {
		return err
	}
	fmt.Printf("%s: %v\r\n", "Name", pokemon.Name)
	fmt.Printf("%s: %v\r\n", "Height", pokemon.Height)
	fmt.Printf("%s: %v\r\n", "Weight", pokemon.Weight)
	fmt.Printf("%s: %v\r\n", "BaseExperience", pokemon.BaseExperience)
	fmt.Println("Pokemon Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("\t- %v: %v\r\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Pokemon types:")
	for _, typ := range pokemon.Types {
		fmt.Printf("\t- %s\r\n", typ.Type.Name)
	}
	fmt.Println("Pokemon Attacks:")
	for _, move := range pokemon.Moves {
		fmt.Printf("\t- %v\r\n", move.Move.Name)
	}
	return nil
}
