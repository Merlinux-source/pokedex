package main

import (
	"boot.dev-Pokedex/internal/pokeapi"
	"encoding/json"
	"fmt"
)

func commandMap(conf *CommandConf) error {
	url := conf.Context["Next"]
	mapApiResponse := pokeapi.NamedAPIResourceList{}
	if url == "" {
		url = pokeapi.APIV2_LocationAreaBaseURL
	}

	resText, err := conf.HttpCache.CacheGet(url)
	if err != nil {
		println(err.Error())
		return err
	}

	json.Unmarshal(resText, &mapApiResponse)

	conf.Context["Next"] = mapApiResponse.Next
	conf.Context["Previous"] = mapApiResponse.Previous

	for _, location := range mapApiResponse.Results {
		fmt.Println(location.Name)
	}

	return nil
}
