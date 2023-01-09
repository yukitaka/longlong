package main

import "github.com/yukitaka/longlong/internal/cmd"

func main() {
	command := cmd.NewLlctlCommand()
	command.Execute()
}
