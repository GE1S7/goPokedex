package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/GE1S7/goPokedex/internal/pokecache"
	"github.com/GE1S7/goPokedex/internal/pokemon"
)

func fetchLocationAreaName(url string, cache *pokecache.Cache) (locationsNames []string, previousUrl string, nextUrl string, err error) {
	var dataJson []byte

	val, ok := cache.Get(url)
	if ok {
		dataJson = val

	} else {
		res, err := http.Get(url)
		if err != nil {
			return []string{}, "", "", fmt.Errorf("API response error")
		}
		dataJson, err = io.ReadAll(res.Body)
		if err != nil {
			return []string{}, "", "", err
		}
		cache.Add(url, dataJson)
	}

	//var data map[string]any
	var data struct {
		Count    float64             `json:"count"`
		Next     string              `json:"next"`
		Previous string              `json:"previous"`
		Results  []map[string]string `json:"results"`
		rest     json.RawMessage
	}

	err = json.Unmarshal(dataJson, &data)
	if err != nil {
		return []string{}, "", "", fmt.Errorf("Error while parsing json")
	}

	for _, e := range data.Results {
		for k, v := range e {
			if k == "name" {
				locationsNames = append(locationsNames, v)
			}
		}
	}

	return locationsNames, data.Previous, data.Next, nil
}

func fetchPokemonData(name string) (pokemonData pokemon.Pokemon, err error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", name)

	res, err := http.Get(url)
	if err != nil {
		return pokemon.Pokemon{}, fmt.Errorf("API response error")
	}

	dataJson, err := io.ReadAll(res.Body)
	if err != nil {
		return pokemon.Pokemon{}, err
	}

	err = json.Unmarshal(dataJson, &pokemonData)
	if err != nil {
		return pokemon.Pokemon{}, fmt.Errorf("Error while parsing json")
	}

	return pokemonData, nil

}
