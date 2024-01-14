package main

import (
	"bufio"
	"os"
	"path/filepath"
)

type CommandHistory struct {
	commands []string
	file     *os.File
}

func (ch *CommandHistory) getHistoryPath() string {
	home, _ := os.UserHomeDir()

	return filepath.Join(home, ".ccsh_history")
}

func (ch *CommandHistory) newHistory() {
	var historyPath = ch.getHistoryPath()
	file, _ := os.OpenFile(historyPath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0600)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var history []string
	for scanner.Scan() {
		history = append(history, scanner.Text())
	}

	ch.commands = history
}

func (ch *CommandHistory) addCommand(input string) {
	ch.commands = append(ch.commands, input)
	ch.file.WriteString(input + "\n")
}

func (ch *CommandHistory) getHistory() []string {
	return ch.commands
}
