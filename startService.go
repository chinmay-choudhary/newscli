package main

import "fmt"

type cliCommand struct {
	name        string
	description string
}

type callback interface {
	Callback(command *string) error
}

func (cmd cliCommand) Callback(command *string) error {
	if cmd.name == "help" {
		err := CommandHelp()
		if err != nil {
			fmt.Println("Error Occured while calling help function please check your input")
		}
		return nil
	} else if cmd.name == "exit" {
		err := CommandExit()
		if err != nil {
			fmt.Println("Count not exit program")
		}
		return nil
	} else if cmd.name == "top" {
		countryCode := &command
		err := CommandTop(**countryCode)
		if err != nil {
			fmt.Println("Could not get top news please check country code")
		}
		return nil
	}
	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
		},
		"top": {
			name:        "top",
			description: "Get Top news of the country",
		},
	}
}
