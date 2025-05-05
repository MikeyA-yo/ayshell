package main

import "strings"

type Completer struct {
	commands []string
}

func NewCompleter() *Completer {
	return &Completer{
		commands: GetCommands(),
	}
}

func (c *Completer) Complete(input string) []string {
	var completions []string
	for _, cmd := range c.commands {
		if strings.HasPrefix(cmd, input) {
			completions = append(completions, cmd)
		}
	}
	return completions
}

func (c *Completer) Suggest(input string) string {
	if input == "" {
		return ""
	}

	minDist := 999
	closest := ""
	for _, cmd := range c.commands {
		dist := levenshtein(input, cmd)
		if dist < minDist {
			minDist = dist
			closest = cmd
		}
	}
	if minDist <= 2 { // Allow small typos
		return closest
	}
	return ""
}

// Simple Levenshtein distance
func levenshtein(a, b string) int {
	la, lb := len(a), len(b)
	dp := make([][]int, la+1)
	for i := range dp {
		dp[i] = make([]int, lb+1)
	}

	for i := 0; i <= la; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= lb; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= la; i++ {
		for j := 1; j <= lb; j++ {
			cost := 0
			if a[i-1] != b[j-1] {
				cost = 1
			}
			dp[i][j] = min(dp[i-1][j]+1, min(dp[i][j-1]+1, dp[i-1][j-1]+cost))
		}
	}

	return dp[la][lb]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
