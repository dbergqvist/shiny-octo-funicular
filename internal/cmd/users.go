package cmd

import "github.com/spf13/cobra"

func newUsersCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "users",
		Short: "Manage Sana users",
	}
}
