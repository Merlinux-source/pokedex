package main

import "fmt"

func commandPokedex(conf *CommandConf) error {
	if len(conf.Pokedex.Pokemon) < 1 {
		fmt.Println("Your pokedex is empty, go catch Jamal!")
		return nil
	}
	fmt.Println("Pokemon in your pokedex:")
	for _, pokemon := range conf.Pokedex.Pokemon {
		fmt.Printf("\t- %s\n", pokemon.Name)
	}
	return nil
}
