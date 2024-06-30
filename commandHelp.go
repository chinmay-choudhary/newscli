package main

import "fmt"

func CommandHelp() error {
	fmt.Println()
	fmt.Println("Welcome To NewsCLI Tool")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
