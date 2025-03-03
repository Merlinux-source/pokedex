package main

import (
	"fmt"
	"sort"
)

func commandHelp(conf *CommandConf) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Commands:")

	// Sort the supportedCommands map by command names
	var commandNames []string
	for name := range supportedCommands {
		commandNames = append(commandNames, name)
	}
	sort.Strings(commandNames)

	for _, name := range commandNames {
		command := supportedCommands[name]
		fmt.Printf("\t%s: %s\r\n", command.name, command.description)
		fmt.Printf("\t\t-> Usage: %s\r\n", command.usage)
	}
	return nil
}
