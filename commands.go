package main

// Command struct remains unchanged
type Command struct {
	Name        string
	Description string
	Execute     func(args []string) error
}

// Remove global executor and commandList initialization
var executor *CommandExecutor
var commandList map[string]Command

func InitCommands() {
	executor = NewCommandExecutor()
	commandList = map[string]Command{
		"cd": {
			Name:        "cd",
			Description: "Change directory",
			Execute:     executor.ChangeDirectory,
		},
		"ls": {
			Name:        "ls",
			Description: "List directory contents",
			Execute:     executor.ListDirectory,
		},
		"echo": {
			Name:        "echo",
			Description: "Print text to console",
			Execute:     executor.Echo,
		},
		"cat": {
			Name:        "cat",
			Description: "Print file contents",
			Execute:     executor.Cat,
		},
		"mkdir": {
			Name:        "mkdir",
			Description: "Create new directory",
			Execute:     executor.Mkdir,
		},
		"rm": {
			Name:        "rm",
			Description: "Remove file or directory",
			Execute:     executor.Rm,
		},
		"touch": {
			Name:        "touch",
			Description: "Create empty file",
			Execute:     executor.Touch,
		},
		"help": {
			Name:        "help",
			Description: "Show available commands",
			Execute:     executor.ShowHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the shell",
			Execute:     executor.ExitShell,
		},
		"mv": {
			Name:        "mv",
			Description: "Move or rename files and directories",
			Execute:     executor.Move,
		},
	}
}

func GetCommands() []string {
	commands := make([]string, 0, len(commandList))
	for k := range commandList {
		commands = append(commands, k)
	}
	return commands
}

func GetCommand(name string) (Command, bool) {
	cmd, exists := commandList[name]
	return cmd, exists
}
