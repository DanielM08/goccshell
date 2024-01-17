# goccshell
Go solution to [Coding Challenges Shell](https://codingchallenges.fyi/challenges/challenge-shell)

## What is a shell ?

Shell is a program that takes commands from the keyboard and gives them to the operating system to perform. In the old days, it was the only user interface available on a Unix-like system such as Linux. Nowadays, we have graphical user interfaces (GUIs) in addition to command line interfaces (CLIs) such as the shell. ([For more details](https://linuxcommand.org/lc3_lts0010.php)) 

## About the goccshell

Our command line shell supports: 

1) Basic individual Unix commands (ls, clear, exit ...).
2) Commands with arguments (diff, ls, mkdir, cat, curl ...).
3) "cd" built-in command to navigate in the system directories.
4) Pipe sequence of commands (e.g., cat test.txt | wc -l).
5) Handle signals that could interrupt the shell (e.g., Ctrl + C).
6) The "history" command.

## Running the project

In the project directory, run ```go run .```

## Project Evolution

1) Add support for running a historic command (ex: allowing the use of the up and down arrows)
2) Support variable expansion. (ex: pressing tab)
3) Handle scripting
