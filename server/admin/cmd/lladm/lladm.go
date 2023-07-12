package main

import (
	"github.com/yukitaka/longlong/server/admin/internal/cmd"
	"log"
)

func main() {
	command := cmd.NewAdminCommand()
	if err := command.Execute(); err != nil {
		log.Fatal(err)
	}
}
