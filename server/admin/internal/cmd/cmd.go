package cmd

import "github.com/spf13/cobra"

func NewAdminCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "admin",
		Short: "LongLong Admin",
		Long:  "LongLong Admin",
	}
}