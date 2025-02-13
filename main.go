package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/aaroosh-07/pokedexcli/internal/pokecache"
	"github.com/aaroosh-07/pokedexcli/internal/pokedex"
)

func main() {
	//create a new scanner
	initCommandRegistry()
	scanner := bufio.NewScanner(os.Stdin)
	var configInfo = &config{
		limit: 20,
		offset: -20,
		cache: pokecache.NewCache(),
		pokedex: pokedex.NewPokedex(),
	}
	for {
		fmt.Print("Pokedex > ")
		//take command input from user
		var command string
		if scanner.Scan() {
			command = scanner.Text()
		}
		tokens := cleanInput(command)
		cmdInfo, isPresent := commandRegistry[tokens[0]]

		if !isPresent {
			fmt.Println("Unknown command")
			continue
		}

		err := cmdInfo.callback(configInfo, tokens[1:])

		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(text)
	var lowercaseWords []string
	for _, word := range words {
		lowercaseWords = append(lowercaseWords, strings.ToLower(word))
	}
	return lowercaseWords
}