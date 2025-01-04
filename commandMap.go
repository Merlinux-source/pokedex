package main

import (
	"boot.dev-Pokedex/internal/pokeapi"
	"encoding/json"
	"fmt"
)

func commandMap(conf *CommandConf) error {
	url := conf.Next
	mapApiResponse := pokeapi.MapApiResponse{}
	if url == "" {
		url = pokeapi.MapBaseUrl
	}

	resText, err := conf.HttpCache.CacheGet(url)
	if err != nil {
		println(err.Error())
		return err
	}

	json.Unmarshal(resText, &mapApiResponse)

	conf.Next = mapApiResponse.Next
	conf.Previous = mapApiResponse.Previous

	for _, location := range mapApiResponse.Results {
		fmt.Println(location.Name)
	}

	return nil
}
