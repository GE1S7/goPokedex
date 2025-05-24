package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/GE1S7/goPokedex/internal/pokecache"
)

func main() {
	interval, err := time.ParseDuration("5s")
	if err != nil {
		fmt.Println(err)
		return
	}
	conf := config{
		previousUrl: "",
		nextUrl:     "https://pokeapi.co/api/v2/location-area",
		cache:       pokecache.NewCache(interval),
	}

	// map of valid commands
	commandRegistry = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the next 20 locations",
			callback:    commandMapb,
		},
		"catch": {
			name:        "catch [name]",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
	}
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print("Pokedex > ")

		var prompt string

		scanner.Scan()

		err := scanner.Err()

		if err != nil {
			log.Fatal(err)
		}

		prompt = scanner.Text()

		if len(prompt) == 0 {
			continue
		}

		args := strings.Split(prompt, " ")

		command, exists := commandRegistry[args[0]]
		if exists {
			err := command.callback(args, &conf)
			if err != nil {
				fmt.Println(err)
			}
		} else {

			fmt.Println("Unknown command")

		}
	}
}

func cleanInput(text string) []string {
	/* split on whitespace, lowercase
	remove leading and trailing whitespace*/

	clean := strings.Split(strings.ToLower(strings.TrimSpace(text)), " ")
	return clean

}
