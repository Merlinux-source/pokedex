package main

import (
	"slices"
	"strings"
)

func cleanInput(input string) []string {
	stripedString := strings.TrimSpace(input)
	newLinesAreSpaces := strings.ReplaceAll(stripedString, "\n", " ")
	splitInput := strings.Split(newLinesAreSpaces, " ")

splitLoop:
	for wordIndex, word := range splitInput {
		splitInput[wordIndex] = strings.ToLower(word)
		if len(word) < 1 {
			splitInput = slices.Delete(splitInput, wordIndex, wordIndex+1)
			goto splitLoop
		}

	}
	return splitInput
}
