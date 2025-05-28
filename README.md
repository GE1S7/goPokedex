# goPokedex REPL

![Pokeball](https://img.icons8.com/color/96/000000/pokeball--v1.png)

A terminal-based Pokédex application that allows you to explore Pokémon locations, catch Pokémon, and manage your collection - all from the command line.

## Features
- Explore Pokémon locations with pagination (`map`/`mapb` commands)
- Discover Pokémon in specific areas (`explore <location>`)
- Catch Pokémon with probability based on experience level (`catch <pokemon>`)
- Inspect caught Pokémon details (stats, types, etc.)
- View your Pokédex collection (`pokedex`)
- Response caching for faster performance
- Simple CLI interface with autocomplete-style commands

## Installation
```bash
go install github.com/GE1S7/goPokedex
```

## Usage
Type `goPokedex` in your terminal to start REPL session


### REPL Commands

```
  map      - Display next 20 locations
  mapb     - Display previous 20 locations
  explore <location> - Find Pokémon in an area
  catch <pokemon>   - Attempt to catch a Pokémon
  inspect <pokemon> - View Pokémon details
  pokedex  - List all caught Pokémon
  help     - Show help
  exit     - Quit the application
```

