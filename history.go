package main

import (
	"bufio"
	"fmt"
	"os"
)

type CommandHistory struct {
	historyFilePath string
}

func (ch *CommandHistory) getHistoryPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return home + "/.ccsh_history"
}

func (ch *CommandHistory) saveHistory(input string) {
	var historyPath = ch.getHistoryPath()
	file, err := os.OpenFile(historyPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(input + "\n")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (ch *CommandHistory) getCommandsHistory() (error, string) {
	var historyPath = ch.getHistoryPath()
	file, err := os.Open(historyPath)
	if err != nil {
		return err, ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var history string
	var prevLine string
	for scanner.Scan() {
		if prevLine != "" {
			history += prevLine + "\n"
		}
		prevLine = scanner.Text()
	}
	history += prevLine

	return nil, history
}
