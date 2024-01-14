package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	for {
		fmt.Print("ccsh> ")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		commandsSequence := strings.Split(input, "|")

		var commandsExecList []*exec.Cmd = []*exec.Cmd{}
		var err error
		for _, commandInput := range commandsSequence {
			err, command := getCommandExecutor(commandInput)

			if err != nil || command == nil {
				break
			}

			commandsExecList = append(commandsExecList, command)
		}

		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}

		err = executePipedCommands(commandsExecList)

		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}
}

func getCommandExecutor(input string) (error, *exec.Cmd) {
	input = strings.TrimSpace(input)
	parts := strings.Fields(input)

	var command = parts[0]
	var args = parts[1:]

	if command == "exit" {
		os.Exit(0)
	}

	if command == "cd" {
		if len(args) == 0 {
			return nil, nil
		}
		dir := args[0]
		if args[0] == "~" {
			dir, _ = os.UserHomeDir()
		}

		err := os.Chdir(dir)
		return err, nil
	}

	cmd := exec.Command(command, args...)

	return nil, cmd
}

func executePipedCommands(commandsExecList []*exec.Cmd) error {

	for i, cmd := range commandsExecList {
		if i > 0 {
			prevCmd := commandsExecList[i-1]
			stdout, err := prevCmd.StdoutPipe()

			if err != nil {
				return err
			}

			cmd.Stdin = stdout
		}

		if i == len(commandsExecList)-1 {
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
		}
	}

	for _, cmd := range commandsExecList {
		err := cmd.Start()

		if err != nil {
			return err
		}
	}

	for _, cmd := range commandsExecList {
		err := cmd.Wait()

		if err != nil {
			if cmd.ProcessState.ExitCode() == -1 {
				return fmt.Errorf("Command not found: %s", cmd.Path)
			}
		}
	}

	return nil
}
