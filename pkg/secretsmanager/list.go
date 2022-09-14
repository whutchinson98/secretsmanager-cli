package secretsmanager

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/spf13/cobra"
)

func ListSecrets(cmd *cobra.Command, args []string) error {
	region, _ := cmd.Flags().GetString("region")

	cfg, _ := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))

	client := secretsmanager.NewFromConfig(cfg)

	res, err := client.ListSecrets(context.TODO(), &secretsmanager.ListSecretsInput{})

	if err != nil {
		return err
	}

	secretNames := make([]string, 0)

	for _, secret := range res.SecretList {
		secretNames = append(secretNames, *secret.Name)
	}

	log.Printf("secrets %v\n", secretNames)

	return nil
}
