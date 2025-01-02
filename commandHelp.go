package main

import "fmt"

func commandHelp() error {

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, c := range supportedCommands {
		fmt.Printf("\t%s: %s\r\n", c.name, c.description)
	}
	return nil
}
