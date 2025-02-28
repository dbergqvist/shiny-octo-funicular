package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sana",
		Short: "Sana CLI - A command line interface for Sana",
		Long:  `A command line tool to interact with the Sana API for managing users, groups, courses, and more.`,
	}

	cmd.PersistentFlags().String("domain", "", "Sana domain (required)")
	cmd.PersistentFlags().String("client-id", "", "OAuth client ID")
	cmd.PersistentFlags().String("client-secret", "", "OAuth client secret")

	viper.BindPFlag("domain", cmd.PersistentFlags().Lookup("domain"))
	viper.BindPFlag("client_id", cmd.PersistentFlags().Lookup("client-id"))
	viper.BindPFlag("client_secret", cmd.PersistentFlags().Lookup("client-secret"))

	cmd.AddCommand(
		newAuthCmd(),
		newUsersCmd(),
		// newGroupsCmd(),
		// newCoursesCmd(),
	)

	return cmd
}
