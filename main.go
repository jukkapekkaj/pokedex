package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
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

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := scanner.Text()
			cleanedInput := cleanInput(input)
			if len(cleanedInput) > 0 {
				command, ok := commands[cleanedInput[0]]
				if ok {
					err := command.callback()
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

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
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
