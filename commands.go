package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/GE1S7/goPokedex/internal/pokecache"
	"github.com/GE1S7/goPokedex/internal/pokemon"
)

type config struct {
	previousUrl string
	nextUrl     string
	cache       *pokecache.Cache
	pokedex     map[string]pokemon.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(args []string, conf *config) error
}

var commandRegistry map[string]cliCommand

func commandHelp(args []string, conf *config) error {
	// display help
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println("")
	for k, v := range commandRegistry {
		fmt.Printf("%s: %s\n", k, v.description)
	}
	return nil
}

func commandMap(args []string, conf *config) error {
	if conf.nextUrl == "" {
		fmt.Println(`You've reached the end of the list. To move back type "mapb".`)
		return nil
	}
	locations, _, next, err := fetchLocationAreaName(conf.nextUrl, conf.cache)
	if err != nil {
		return err
	} else {
		conf.previousUrl = conf.nextUrl
		conf.nextUrl = next
	}

	for _, e := range locations {
		fmt.Println(e)
	}

	return nil

}

func commandMapb(args []string, conf *config) error {
	if conf.previousUrl == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	locations, previous, _, err := fetchLocationAreaName(conf.previousUrl, conf.cache)

	if err != nil {
		return err
	} else {
		conf.nextUrl = conf.previousUrl
		conf.previousUrl = previous
	}

	for _, e := range locations {
		fmt.Println(e)
	}

	return nil
}

func commandCatch(args []string, conf *config) error {
	name := args[1]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	/*
		use pokemon endpoint (GET https://pokeapi.co/api/v2/pokemon/{id or name}/)
		use math/rand package to make success dependent on chance
		use "base_experience" field (higher/harder)
		add to pokedex map[string]Pokemon

	*/
	baseExperience, err := fetchPokemonData(name)
	if err != nil {
		return err
	} else {
		if rand.Intn(baseExperience) < 11 {
			conf.pokedex = append(conf.pokedex, name)
			fmt.Println(name, "was caught!")
		} else {
			fmt.Println(name, "escaped!")
		}
	}
	return nil
}

func commandExit(args []string, conf *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
