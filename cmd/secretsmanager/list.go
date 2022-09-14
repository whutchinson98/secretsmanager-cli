package secretsmanager

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/whutchinson98/secretsmanager-cli/pkg/secretsmanager"
)

var ListCmd = &cobra.Command{
	Use:   "list-secrets",
	Short: "Lists secrets in SecretsManager",
	Long:  `Lists all secrets of a given region in SecretsManager.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := secretsmanager.ListSecrets(cmd, args)
		if err != nil {
			fmt.Printf("Error listing secrets %s", err.Error())
		}
	},
}

func InitList() {
	ListCmd.Flags().StringP("region", "r", "us-east-1", "The AWS region for the secrets you want to see")
}
