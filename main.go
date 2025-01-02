package main

import (
	"bufio"
	"fmt"
	"os"
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
			conf:        &CommandConf{},
		},
		"mapb": cliCommand{
			name:        "mapb",
			description: "Paginate the Map locations backwards",
			callback:    CommandMapb,
			conf:        &CommandConf{},
		},
	}
}

func main() {
	buf := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		buf.Scan()
		cleanInput := cleanInput(buf.Text())

		command, ok := supportedCommands[cleanInput[0]]
		if !ok {
			fmt.Printf("Unknown command: %s\n", cleanInput[0])
		}
		if ok {
			err := command.callback(command.conf)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
			}
		}
	}
}
