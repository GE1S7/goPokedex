package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	conf := config{
		previousUrl: "",
		nextUrl:     "https://pokeapi.co/api/v2/location-area",
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

		command, exists := commandRegistry[prompt]
		if exists {
			err := command.callback(&conf)
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
