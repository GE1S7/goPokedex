package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

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

		fmt.Println("Your command was:", cleanInput(prompt)[0])

	}
}

func cleanInput(text string) []string {
	/* split on whitespace, lowercase
	remove leading and trailing whitespace*/

	v := strings.Split(strings.ToLower(strings.TrimSpace(text)), " ")
	return v

}
