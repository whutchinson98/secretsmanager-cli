package secretsmanager

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/whutchinson98/secretsmanager-cli/pkg/secretsmanager"
)

var RemoveCmd = &cobra.Command{
	Use:   "remove-secret",
	Short: "Removes a SecretsManager secret",
	Long:  `Allows you to easily remove a secret from SecretsManager.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := secretsmanager.RemoveSecret(cmd, args)
		if err != nil {
			fmt.Printf("Error listing secrets %s", err.Error())
		}
	},
}

func InitRemove() {
	RemoveCmd.Flags().StringP("region", "r", "us-east-1", "The AWS region the Secret is located")
	RemoveCmd.Flags().StringP("secret-name", "s", "", "The Secret name (Must be specified)")
}
