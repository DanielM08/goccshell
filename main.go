package main

func main() {
	ch := CommandHistory{}
	shell := Shell{&ch}

	shell.execute()
}
