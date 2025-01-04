package main

import "boot.dev-Pokedex/internal/pokecache"

type CommandConf struct {
	Next      string
	Previous  string
	HttpCache *pokecache.Cache
}
