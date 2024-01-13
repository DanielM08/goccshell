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

		parts := strings.Fields(input)
		var command = parts[0]
		var args = parts[1:]

		if command == "exit" {
			os.Exit(0)
		}

		if command == "cd" {
			if len(args) == 0 {
				continue
			}

			if args[0] == "~" {
				dir, _ := os.UserHomeDir()
				err := os.Chdir(dir)

				if err != nil {
					fmt.Printf("%v\n", err)
				}

				continue
			}

			err := os.Chdir(args[0])

			if err != nil {
				fmt.Printf("%v\n", err)
			}

			continue
		}

		cmd := exec.Command(command, args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()

		if err != nil {
			fmt.Println(err)
		}
	}
}
