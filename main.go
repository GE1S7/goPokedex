package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// map of valid commands
var commandRegistry = map[string]cliCommand{
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
}

func main() {
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
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Unknown command")
			}
		}

	}
}

func commandHelp() error {
	return nil
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func cleanInput(text string) []string {
	/* split on whitespace, lowercase
	remove leading and trailing whitespace*/

	clean := strings.Split(strings.ToLower(strings.TrimSpace(text)), " ")
	return clean

}
