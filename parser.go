package main

import (
	"strings"
	"unicode"
)

// parseCommandWithQuotes parses a command string into arguments,
// respecting quoted strings as single arguments
func parseCommandWithQuotes(input string) []string {
	var args []string
	var currentArg strings.Builder
	inQuotes := false
	quoteChar := '"' // Default to double quotes

	// Trim leading/trailing whitespace
	input = strings.TrimSpace(input)

	// Handle empty input
	if input == "" {
		return args
	}

	for _, char := range input {
		// Handle quotes (both single and double)
		if char == '"' || char == '\'' {
			if !inQuotes {
				// Starting a quoted section
				inQuotes = true
				quoteChar = char
			} else if char == quoteChar {
				// Ending a quoted section if quote type matches
				inQuotes = false
			} else {
				// Different quote character inside quotes, treat as normal character
				currentArg.WriteRune(char)
			}
			continue
		}

		// Handle spaces
		if unicode.IsSpace(char) && !inQuotes {
			// Space outside quotes means end of current argument
			if currentArg.Len() > 0 {
				args = append(args, currentArg.String())
				currentArg.Reset()
			}
			continue
		}

		// Add character to current argument
		currentArg.WriteRune(char)
	}

	// Add the last argument if there is one
	if currentArg.Len() > 0 {
		args = append(args, currentArg.String())
	}

	return args
}
