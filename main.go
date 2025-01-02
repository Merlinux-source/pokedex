package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		buf.Scan()
		cleanInput := cleanInput(buf.Text())

	}
}
