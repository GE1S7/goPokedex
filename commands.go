package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/GE1S7/goPokedex/internal/pokecache"
	"github.com/GE1S7/goPokedex/internal/pokemon"
)

type config struct {
	previousUrl  string
	nextUrl      string
	mapCache     *pokecache.Cache
	pokemonCache *pokecache.Cache
	pokedex      map[string]pokemon.Pokemon
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
	locations, _, next, err := fetchLocationAreaName(conf.nextUrl, conf.mapCache)
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

	locations, previous, _, err := fetchLocationAreaName(conf.previousUrl, conf.mapCache)

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

func commandExplore(args []string, conf *config) error {

	pokemonFound, err := fetchLocationAreaPokemonData(args[1], conf.pokemonCache)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", args[1])

	for _, e := range pokemonFound {
		fmt.Println(e)
	}

	return nil

}

func commandCatch(args []string, conf *config) error {
	name := args[1]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	pokemonData, err := fetchPokemonData(name)
	if err != nil {
		return err
	} else {
		if rand.Intn(pokemonData.BaseExperience) < 11 {
			conf.pokedex[name] = pokemonData
			fmt.Println(name, "was caught!")
		} else {
			fmt.Println(name, "escaped!")
		}
	}
	return nil
}

func commandInspect(args []string, conf *config) error {
	_, ok := conf.pokedex[args[1]]
	if ok {
		pokemonData, err := fetchPokemonData(args[1])
		if err != nil {
			return err
		}
		fmt.Println("Name:", pokemonData.Name)
		fmt.Println("Height:", pokemonData.Height)
		fmt.Println("Weight:", pokemonData.Weight)
		fmt.Println("Stats:")
		for _, e := range pokemonData.Stats {
			fmt.Printf("  -%s: %d\n", e.Stat.Name, e.BaseStat)
		}
		fmt.Println("Stats:")
		for _, e := range pokemonData.Types {
			fmt.Printf("  -%s\n", e.Type.Name)
		}

	} else {
		fmt.Println("you have not caught that pokemon")
	}
	return nil
}

func commandPokedex(args []string, conf *config) error {
	fmt.Println("Your Pokedex:")
	for k := range conf.pokedex {
		fmt.Printf("  -%s\n", k)
	}
	return nil
}

func commandExit(args []string, conf *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
