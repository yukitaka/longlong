package main

import (
	"fmt"
	"github.com/yukitaka/longlong/server/admin/internal/cmd"
)

func main() {
	command := cmd.NewAdminCommand()
	fmt.Printf("%#v", command)
}
