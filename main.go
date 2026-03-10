package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := scanner.Text()
			lower := strings.ToLower(input)
			words := strings.Fields(lower)
			if len(words) > 0 {
				firstWord := words[0]
				fmt.Println("Your command was:", firstWord)
			} else {
				fmt.Println("No command given.")
			}
		}
	}
}
