package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/peterh/liner"
)

func main() {
	line := liner.NewLiner()
	defer line.Close()
	InitCommands()
	executor.SetBuiltins(commandList)

	completer := NewCompleter()

	// Autocompletion setup
	line.SetCompleter(func(line string) []string {
		return completer.Complete(line)
	})

	fmt.Printf(`
	 █████╗ ██╗   ██╗    ███████╗██╗  ██╗███████╗██╗     ██╗     
	██╔══██╗╚██╗ ██╔╝    ██╔════╝██║  ██║██╔════╝██║     ██║     
	███████║ ╚████╔╝     ███████╗███████║█████╗  ██║     ██║     
	██╔══██║  ╚██╔╝      ╚════██║██╔══██║██╔══╝  ██║     ██║     
	██║  ██║   ██║       ███████║██║  ██║███████╗███████╗███████╗
	╚═╝  ╚═╝   ╚═╝       ╚══════╝╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝
	Type 'help' for commands or 'exit' to quit
`)

	for {
		curDir, err := os.Getwd()
		if err != nil {
			fmt.Println("Error getting current directory:", err)
			return
		}
		fmt.Println()
		input, err := line.Prompt(fmt.Sprintf("aysh %s> ", curDir))
		fmt.Println()
		if err != nil {
			if err == liner.ErrPromptAborted {
				fmt.Println("Aborted")
				break
			}
			fmt.Println("Error reading line:", err)
			break
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		line.AppendHistory(input)

		if input == "exit" {
			break
		}

		// Use the executor to handle both built-in and system commands
		if err := executor.Execute(input); err != nil {
			// If command not found, try to suggest a similar built-in command
			cmdName := strings.Fields(input)[0]
			suggestion := completer.Suggest(cmdName)
			fmt.Printf("Error: %v\n", err)
			if suggestion != "" {
				fmt.Printf("Did you mean: '%s'?\n", suggestion)
			}
		}
	}
}
