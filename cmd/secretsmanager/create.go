package secretsmanager

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/whutchinson98/secretsmanager-cli/pkg/secretsmanager"
)

var CreateCmd = &cobra.Command{
	Use:   "create-secret",
	Short: "Creates a SecretsManager secret given an env file",
	Long:  `Allows you to easily take down a large .env file and convert it to an SecretsManager secret.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := secretsmanager.CreateSecret(cmd, args)
		if err != nil {
			fmt.Printf("Error creating secret %s", err.Error())
		}
	},
}

func InitCreate() {
	CreateCmd.Flags().StringP("region", "r", "us-east-1", "The AWS region the Secret is located")
	CreateCmd.Flags().StringP("secret-name", "s", "", "The Secret name (Must be specified)")
	CreateCmd.Flags().StringP("env-file", "e", ".env", "The name of the env file")
	CreateCmd.Flags().BoolP("new-secret", "n", false, "If enabled this will create a new Secret in aws")
}
