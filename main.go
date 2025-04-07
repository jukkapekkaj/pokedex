package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
	strs := "   Hello    World    "
	formatted_strings := make([]string, 0)
	for _, s := range strings.Split(strs, " ") {
		if len(s) > 0 && s != " " {
			formatted_strings = append(formatted_strings, strings.TrimSpace(s))
		}
	}

	fmt.Println(formatted_strings)
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
