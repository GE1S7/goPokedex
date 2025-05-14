package main

import (
	"fmt"
	"os"
)

type config struct {
	previousUrl string
	nextUrl     string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

var commandRegistry map[string]cliCommand

func commandHelp(conf *config) error {
	// display help
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println("")
	for k, v := range commandRegistry {
		fmt.Printf("%s: %s\n", k, v.description)
	}
	return nil
}

func commandMap(conf *config) error {
	if conf.nextUrl == "" {
		fmt.Println(`You've reached the end of the list. To move back type "mapb".`)
		return nil
	}
	locations, _, next, err := fetchLocationAreaName(conf.nextUrl)
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

func commandMapb(conf *config) error {
	if conf.previousUrl == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	locations, previous, _, err := fetchLocationAreaName(conf.previousUrl)

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

func commandExit(conf *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
