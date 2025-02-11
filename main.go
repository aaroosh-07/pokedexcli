package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/aaroosh-07/pokedexcli/internal/pokecache"
)

func main() {
	//create a new scanner
	scanner := bufio.NewScanner(os.Stdin)
	var configInfo = &config{
		limit: 20,
		offset: -20,
		cache: pokecache.NewCache(),
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

		cmdInfo.callback(configInfo)
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