package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yourusername/sana-cli/internal/api"
)

func newAuthCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Authenticate with Sana API",
		RunE: func(cmd *cobra.Command, args []string) error {
			domain := viper.GetString("domain")
			clientID := viper.GetString("client_id")
			clientSecret := viper.GetString("client_secret")

			if domain == "" || clientID == "" || clientSecret == "" {
				return fmt.Errorf("domain, client-id, and client-secret are required")
			}

			client := api.NewClient(domain)
			token, err := client.GetToken(clientID, clientSecret)
			if err != nil {
				return fmt.Errorf("getting token: %w", err)
			}

			fmt.Printf("Access Token: %s\n", token.Data.AccessToken)
			fmt.Printf("Token Type: %s\n", token.Data.TokenType)
			fmt.Printf("Expires In: %d seconds\n", token.Data.ExpiresIn)

			// Save token to config file
			viper.Set("access_token", token.Data.AccessToken)
			if err := viper.WriteConfig(); err != nil {
				return fmt.Errorf("saving token to config: %w", err)
			}

			return nil
		},
	}

	return cmd
}
