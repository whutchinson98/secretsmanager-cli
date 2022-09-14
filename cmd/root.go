package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/whutchinson98/secretsmanager-cli/cmd/env"
	"github.com/whutchinson98/secretsmanager-cli/cmd/secretsmanager"
)

var rootCmd = &cobra.Command{
	Use:   "secretsmanager-to-env",
	Short: "Utilities to work with AWS SecretsManager and Env Files",
	Long: `This package is used to quickly pull down a JSON secret from AWS and create a .env file for it for your
         projects that do not use the secrets via aws-sdk`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
	rootCmd.AddCommand(env.CreateCmd)
	rootCmd.AddCommand(secretsmanager.CreateCmd)
	rootCmd.AddCommand(secretsmanager.ListCmd)
	rootCmd.AddCommand(secretsmanager.RemoveCmd)
}
