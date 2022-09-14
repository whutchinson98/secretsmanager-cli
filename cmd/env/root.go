package env

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/whutchinson98/secretsmanager-cli/pkg/env"
)

var CreateCmd = &cobra.Command{
	Use:   "create-env",
	Short: "Creates an environment variable file from a provided secret",
	Long:  `Allows you to easily take down a large secret from secretsmanager and convert it to a .env file.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := env.CreateEnvFile(cmd, args)
		if err != nil {
			fmt.Printf("Error creating env file %s", err.Error())
		}
	},
}

func init() {
	CreateCmd.Flags().StringP("region", "r", "us-east-1", "The AWS region the Secret is located")
	CreateCmd.Flags().StringP("secret-name", "s", "", "The Secret name (Must be specified)")
	CreateCmd.Flags().StringP("env-file", "e", ".env", "The name of the env file")
	CreateCmd.Flags().StringP("path", "p", "./", "The relative path to place the env file in")
}
