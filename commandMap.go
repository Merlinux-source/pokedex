package main

import (
	"boot.dev-Pokedex/internal/pokeapi"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func commandMap(conf *CommandConf) error {
	url := conf.Next
	if url == "" {
		url = pokeapi.MapBaseUrl
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	mapApiResponse := pokeapi.MapApiResponse{}
	resText, err := ioutil.ReadAll(res.Body)
	if err != nil {
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
