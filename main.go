package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Newscli >: ")
		var CountryCode string
		reader.Scan()

		userInput := reader.Text()

		words := formatUserInput(userInput)

		commandName := words[0]

		command, exists := getCommands()[commandName]

		if !exists {
			fmt.Println("Command not found please enter help tp get the list of commands supported")
		}
		CountryCode, err := getElement(words, 1)
		if err != nil {
		}
		command.Callback(&CountryCode)

	}
}

func formatUserInput(input string) []string {
	lowerInput := strings.ToLower(input)
	splitInput := strings.Fields(lowerInput)
	return splitInput
}

func getElement(arr []string, index int) (string, error) {
	if index < 0 || index >= len(arr) {
		return "", fmt.Errorf("index out of bounds")
	}
	return arr[index], nil
}
