package main

import "boot.dev-Pokedex/internal/pokecache"

type CommandConf struct {
	Context   map[string]string
	HttpCache *pokecache.Cache
	Argv      []string
}
