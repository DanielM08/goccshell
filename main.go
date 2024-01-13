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

		commandsList := strings.Split(input, "|")

		var commandsExecList []*exec.Cmd = []*exec.Cmd{}
		var err error
		for _, commandInput := range commandsList {
			err, command := executeCommand(commandInput)

			if err != nil {
				break
			}

			if command == nil {
				break
			}

			commandsExecList = append(commandsExecList, command)
		}

		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}

		if len(commandsExecList) == 1 {
			cmd := commandsExecList[0]
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			err = cmd.Run()

		} else {
			executePipedCommands(commandsExecList)
		}

		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}
}

func executeCommand(input string) (error, *exec.Cmd) {
	parts := strings.Fields(input)

	var err error

	var command = parts[0]
	var args = parts[1:]

	if command == "exit" {
		os.Exit(0)
	}

	if command == "cd" {
		if len(args) == 0 {
			return nil, nil
		}

		if args[0] == "~" {
			dir, _ := os.UserHomeDir()
			err = os.Chdir(dir)

			if err != nil {
				return nil, nil
			}

			return nil, nil
		}

		err := os.Chdir(args[0])

		if err != nil {
			return err, nil
		}

		return nil, nil
	}

	cmd := exec.Command(command, args...)
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr

	// err = cmd.Run()

	// if err != nil {
	// 	return err, nil
	// }

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
			return err
		}
	}

	return nil
}
