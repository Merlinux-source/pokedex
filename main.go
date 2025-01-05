package main

import (
	"boot.dev-Pokedex/internal/pokeapi"
	"boot.dev-Pokedex/internal/pokecache"
	"bufio"
	"errors"
	"fmt"
	"os"
	"time"
)

var supportedCommands map[string]cliCommand

func init() {
	supportedCommands = map[string]cliCommand{
		"exit": cliCommand{
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
			conf:        &CommandConf{},
			usage:       "exit",
		},
		"help": cliCommand{
			name:        "help",
			description: "Show this help message",
			callback:    commandHelp,
			conf:        &CommandConf{},
			usage:       "help",
		},
		"map": cliCommand{
			name:        "map",
			description: "Paginate the Map locations forward",
			callback:    commandMap,
			conf:        &CommandConf{Context: map[string]string{"Next": "", "Previous": ""}},
			usage:       "map",
		},
		"explore": cliCommand{
			name:        "explore",
			description: "Explore the Pokemon available at location X",
			callback:    commandExplore,
			conf:        &CommandConf{},
			usage:       "explore [Location]",
		},
		"catch": cliCommand{
			name:        "catch",
			description: "Catch the Pokemon",
			callback:    commandCatch,
			conf:        &CommandConf{},
			usage:       "catch [Pokemon]",
		},
		"inspect": cliCommand{
			name:        "inspect",
			description: "Inspect the Pokemon",
			callback:    commandInspect,
			conf:        &CommandConf{},
			usage:       "inspect [Pokemon]",
		},
	}
	supportedCommands["mapb"] = cliCommand{ // this has to be defined separately so that the map navigation share the next and the previous url.
		name:        "mapb",
		description: "Paginate the Map locations backwards",
		callback:    CommandMapb,
		conf:        supportedCommands["map"].conf,
		usage:       "mapb",
	}
}

type Pokedex struct {
	Pokemon []pokeapi.Pokemon
	Count   int
}

func (pd *Pokedex) Add(pkmn pokeapi.Pokemon) {
	pd.Pokemon = append(pd.Pokemon, pkmn)
}

func (pd *Pokedex) Get(name string) (pokeapi.Pokemon, error) {
	for _, pok := range pd.Pokemon {
		if pok.Name == name {
			return pok, nil
		}
	}
	return pokeapi.Pokemon{}, errors.New("Pokedex does not contain the given name")
}

func main() {
	buf := bufio.NewScanner(os.Stdin)
	httpCache := pokecache.NewCache(10 * time.Second)
	pokedex := Pokedex{}
	for _, cmd := range supportedCommands { // insert the http cache reference to the command configuration.
		cmd.conf.HttpCache = httpCache
		cmd.conf.Pokedex = &pokedex
	}

	for {
		fmt.Print("Pokedex > ")
		buf.Scan()
		cleanInput := cleanInput(buf.Text())

		command, ok := supportedCommands[cleanInput[0]]
		if !ok {
			fmt.Printf("Unknown command: %s\n", cleanInput[0])
		}
		if ok {
			command.conf.Argv = cleanInput
			err := command.callback(command.conf)
			if err != nil {
				fmt.Printf("A error occured calling %s Error: %v\n", cleanInput[0], err)
			}
		}
	}
}
