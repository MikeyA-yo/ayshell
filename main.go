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
	line.SetCtrlCAborts(true)

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

		cmdName := strings.Fields(input)[0]
		if cmd, exists := GetCommand(cmdName); exists {
			if cmd.Execute != nil {
				args := strings.Fields(input)[1:]
				if err := cmd.Execute(args); err != nil {
					fmt.Printf("Error: %v\n", err)
				}
			} else {
				fmt.Printf("Command '%s' not implemented yet\n", cmdName)
			}
		} else {
			suggestion := completer.Suggest(cmdName)
			fmt.Printf("Unknown command: '%s'\n", cmdName)
			if suggestion != "" {
				fmt.Printf("Did you mean: '%s'?\n", suggestion)
			}
		}
	}
}
