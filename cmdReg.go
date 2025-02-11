package main

import (
	"fmt"
	"os"

	"github.com/aaroosh-07/pokedexcli/internal/pokeapi"
	"github.com/aaroosh-07/pokedexcli/internal/pokecache"
)

type config struct {
	limit int
	offset int
	cache *pokecache.Cache
}

type cliCommand struct {
	name string
	description string
	callback func(*config, []string) error
}

var commandRegistry map[string]cliCommand

func initCommandRegistry() {
	commandRegistry = map[string] cliCommand {
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "displays basic command info",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "prints the next 20 locations in pokemon world",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "prints the previous 20 locations in pokemon world",
			callback: commandMapb,
		},
		"explore": {
			name: "explore",
			description: "explores a particular location in pokemon world and return pokemon names",
			callback: commandExplore,
		},
	}
}

func commandExit(_ *config, _ []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(_ *config, _ []string) error {
	helpstr := "Welcome to the Pokedex!\nUsage:\n"

	for _, val := range commandRegistry {
		cmdDes := fmt.Sprintf("%s: %s", val.name, val.description)
		helpstr = fmt.Sprintf("%s\n%s",helpstr, cmdDes)
	}
	fmt.Println(helpstr)
	return nil
}

func commandMap(c *config, _ []string) error {
	c.offset += 20
	pokiLociInfo, err := pokeapi.GetPokeapiLocation(c.limit, c.offset, c.cache)

	if err != nil {
		return err
	}

	for _, location := range pokiLociInfo.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(c *config, _ []string) error {
	if c.offset <= 0 {
		c.offset = 0
	} else {
		c.offset -= 20
	}

	pokiLociInfo, err := pokeapi.GetPokeapiLocation(c.limit, c.offset, c.cache)

	if err != nil {
		return err
	}

	for _, location := range pokiLociInfo.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandExplore(c *config, tokens []string) error {
	if len(tokens) < 1 {
		return fmt.Errorf("location name needed for explore cmd")
	}

	pokemonList, err := pokeapi.GetPokemonAtLocation(tokens[0], c.cache)

	if err != nil {
		return err
	}

	for _, pokename := range pokemonList {
		fmt.Println(pokename)
	}

	return nil
}