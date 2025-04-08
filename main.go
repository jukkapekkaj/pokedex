package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jukkapekkaj/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(c *pokeapi.Config) error
}

var commands = map[string]cliCommand{}

func main() {
	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}

	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}

	commands["map"] = cliCommand{
		name:        "map",
		description: "Show next locations",
		callback:    pokeapi.GetNextMap,
	}

	commands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Show previous locations",
		callback:    pokeapi.GetPrevMap,
	}

	poke_config := &pokeapi.Config{Next: "", Previous: ""}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\nPokedex > ")
		if scanner.Scan() {
			input := scanner.Text()
			cleanedInput := cleanInput(input)
			if len(cleanedInput) > 0 {
				command, ok := commands[cleanedInput[0]]
				if ok {
					err := command.callback(poke_config)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println("Unknown command")
				}
			}
		} else {
			fmt.Println("Scanner return false")
		}

	}

}

func commandExit(c *pokeapi.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *pokeapi.Config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println()
	for name, command := range commands {
		fmt.Printf("%s: %s\n", name, command.description)
	}
	return nil
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	formatted_strings := make([]string, 0)
	for _, s := range strings.Split(text, " ") {
		if len(s) > 0 && s != " " {
			formatted_strings = append(formatted_strings, strings.TrimSpace(s))
		}
	}
	return formatted_strings
}
