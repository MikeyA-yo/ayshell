package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type CommandExecutor struct {
	builtins map[string]Command
	shell    string // cmd.exe for Windows
}

func NewCommandExecutor() *CommandExecutor {
	executor := &CommandExecutor{
		builtins: nil, // Will be set after InitCommands
		shell:    "cmd.exe",
	}
	return executor
}

// Remove initializeCommands and builtins assignment from constructor

func (e *CommandExecutor) SetBuiltins(cmds map[string]Command) {
	e.builtins = cmds
}

func (e *CommandExecutor) Execute(input string) error {
	// Parse input respecting quotes
	args := parseCommandWithQuotes(input)
	if len(args) == 0 {
		return nil
	}

	// Check for built-in commands first
	if e.builtins != nil {
		if cmd, exists := e.builtins[args[0]]; exists {
			if cmd.Execute != nil {
				return cmd.Execute(args[1:])
			}
			return fmt.Errorf("command '%s' not implemented yet", args[0])
		}
	}

	// Try to execute as system command
	return e.executeSystem(args)
}

func (e *CommandExecutor) executeSystem(args []string) error {
	// Check if the command exists in PATH
	cmdPath, err := exec.LookPath(args[0])
	ayshPrint := `
     █████╗ ██╗   ██╗    ███████╗██╗  ██╗███████╗██╗     ██╗     
	██╔══██╗╚██╗ ██╔╝    ██╔════╝██║  ██║██╔════╝██║     ██║     
	███████║ ╚████╔╝     ███████╗███████║█████╗  ██║     ██║     
	██╔══██║  ╚██╔╝      ╚════██║██╔══██║██╔══╝  ██║     ██║     
	██║  ██║   ██║       ███████║██║  ██║███████╗███████╗███████╗
	╚═╝  ╚═╝   ╚═╝       ╚══════╝╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝
	`
	if err != nil {
		return fmt.Errorf("%v\ncommand not found: %s", ayshPrint, args[0])
	}

	cmd := exec.Command(cmdPath, args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Set current directory as working directory
	cmd.Dir, err = os.Getwd()
	if err != nil {
		return err
	}

	return cmd.Run()
}

// Helper functions for built-in commands
func (e *CommandExecutor) ChangeDirectory(args []string) error {
	if len(args) == 0 {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		return os.Chdir(homeDir)
	}

	path := args[0]
	if !filepath.IsAbs(path) {
		current, err := os.Getwd()
		if err != nil {
			return err
		}
		path = filepath.Join(current, path)
	}

	return os.Chdir(path)
}

func (e *CommandExecutor) ListDirectory(args []string) error {
	path := "."
	if len(args) > 0 {
		path = args[0]
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue
		}
		if entry.IsDir() {
			fmt.Printf("[DIR] %s\n", entry.Name())
		} else {
			fmt.Printf("[%8d] %s\n", info.Size(), entry.Name())
		}
	}
	return nil
}

func (e *CommandExecutor) ShowHelp(args []string) error {
	fmt.Printf(`
	Welcome to AY SHELL!
	 █████╗ ██╗   ██╗    ███████╗██╗  ██╗███████╗██╗     ██╗     
	██╔══██╗╚██╗ ██╔╝    ██╔════╝██║  ██║██╔════╝██║     ██║     
	███████║ ╚████╔╝     ███████╗███████║█████╗  ██║     ██║     
	██╔══██║  ╚██╔╝      ╚════██║██╔══██║██╔══╝  ██║     ██║     
	██║  ██║   ██║       ███████║██║  ██║███████╗███████╗███████╗
	╚═╝  ╚═╝   ╚═╝       ╚══════╝╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝
	Built with ❤️  by AY (Mikey)
	`)
	fmt.Println("Available commands:")
	for _, cmd := range e.builtins {
		fmt.Printf("%s - %s\n", cmd.Name, cmd.Description)
	}
	return nil
}

func (e *CommandExecutor) ExitShell(args []string) error {
	fmt.Println("Exiting AY SHELL...")
	os.Exit(0)
	return nil
}

func (e *CommandExecutor) Touch(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("touch: missing file operand")
	}

	for _, file := range args {
		_, err := os.Create(file)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *CommandExecutor) Cat(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("cat: missing file operand")
	}

	for _, file := range args {
		content, err := os.ReadFile(file)
		if err != nil {
			return err
		}
		fmt.Println(string(content))
	}
	return nil
}

func (e *CommandExecutor) Echo(args []string) error {
	fmt.Println(strings.Join(args, " "))
	return nil
}

func (e *CommandExecutor) Clear(args []string) error {
	cmd := exec.Command(e.shell, "/c", "cls")
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func (e *CommandExecutor) Mkdir(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("mkdir: missing directory operand")
	}

	for _, dir := range args {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *CommandExecutor) Move(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("mv: missing destination file operand after '%s'", args[0])
	}
	source := args[0]
	destination := args[1]

	err := os.Rename(source, destination)
	if err != nil {
		return err
	}
	return nil
}

func (e *CommandExecutor) Rm(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("rm: missing file operand")
	}

	for _, file := range args {
		err := os.Remove(file)
		if err != nil {
			return err
		}
	}
	return nil
}
