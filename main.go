package main

import (
	"boot.dev-Pokedex/internal/pokecache"
	"bufio"
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
		},
		"help": cliCommand{
			name:        "help",
			description: "Show this help message",
			callback:    commandHelp,
			conf:        &CommandConf{},
		},
		"map": cliCommand{
			name:        "map",
			description: "Paginate the Map locations forward",
			callback:    commandMap,
			conf:        &CommandConf{Context: map[string]string{"Next": "", "Previous": ""}},
		},
		"explore": cliCommand{
			name:        "explore",
			description: "Explore the Pokemon available at location X",
			callback:    commandExplore,
			conf:        &CommandConf{},
		},
	}
	supportedCommands["mapb"] = cliCommand{ // this has to be defined seperately so that the map navigation share the next and the previous url.
		name:        "mapb",
		description: "Paginate the Map locations backwards",
		callback:    CommandMapb,
		conf:        supportedCommands["map"].conf,
	}
}

func main() {
	buf := bufio.NewScanner(os.Stdin)
	httpCache := pokecache.NewCache(10 * time.Second)
	for _, cmd := range supportedCommands { // insert the http cache reference to the command configuration.
		cmd.conf.HttpCache = httpCache
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
				fmt.Printf("Error: %s\n", err)
			}
		}
	}
}
