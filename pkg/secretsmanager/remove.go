package secretsmanager

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/spf13/cobra"
)

func RemoveSecret(cmd *cobra.Command, args []string) error {
	region, _ := cmd.Flags().GetString("region")

	cfg, _ := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))

	client := secretsmanager.NewFromConfig(cfg)

	secretName, _ := cmd.Flags().GetString("secret-name")

	_, err := client.DeleteSecret(context.TODO(), &secretsmanager.DeleteSecretInput{
		SecretId: &secretName,
	})

	if err != nil {
		return fmt.Errorf("removing secret %v", err)
	}

	return nil
}
