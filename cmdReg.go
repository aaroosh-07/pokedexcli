package main

import (
	"fmt"
	"os"

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
	callback func(*config) error
}

var commandRegistry = map[string] cliCommand {
	"exit": {
		name: "exit",
		description: "Exit the Pokedex",
		callback: commandExit,
	},
	"help": {
		name: "help",
		description: "prints documents for using help message",
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
}

func commandExit(_ *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(_ *config) error {
	helpstr := 
`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex`
	fmt.Println(helpstr)
	return nil
}

func commandMap(c *config) error {
	c.offset += 20
	pokiLociInfo, err := getPokeapiLocation(c.limit, c.offset, c.cache)

	if err != nil {
		return err
	}

	for _, location := range pokiLociInfo.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(c *config) error {
	if c.offset <= 0 {
		c.offset = 0
	} else {
		c.offset -= 20
	}

	pokiLociInfo, err := getPokeapiLocation(c.limit, c.offset, c.cache)

	if err != nil {
		return err
	}

	for _, location := range pokiLociInfo.Results {
		fmt.Println(location.Name)
	}

	return nil
}