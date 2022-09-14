package secretsmanager

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/spf13/cobra"
	"github.com/whutchinson98/secretsmanager-cli/internal/util"
)

func CreateSecret(cmd *cobra.Command, args []string) error {
	region, _ := cmd.Flags().GetString("region")

	cfg, _ := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))

	client := secretsmanager.NewFromConfig(cfg)

	secretName, _ := cmd.Flags().GetString("secret-name")

	envFilePath, _ := cmd.Flags().GetString("env-file")

	workingDir, _ := os.Getwd()

	file, err := os.Open(workingDir + "/" + envFilePath)

	if err != nil {
		return fmt.Errorf("opening file: %v", err)
	}

	var perline string
	envFileData := make([]string, 0)

	for {
		_, err := fmt.Fscanf(file, "%v\n", &perline) // give a pattern to scan

		if err != nil {
			if err == io.EOF {
				break // stop reading the file
			}
			return fmt.Errorf("reading file: %v", err)
		}

		envFileData = append(envFileData, perline)
	}

	stringEnvFile, err := util.BuildJsonStringFromEnv(envFileData)

	if err != nil {
		return fmt.Errorf("unable to construct json string from env %v", err.Error())
	}

	newSecret, _ := cmd.Flags().GetBool("new-secret")

	if newSecret {
		log.Printf("Creating secret %s", secretName)
		_, err := client.CreateSecret(context.TODO(), &secretsmanager.CreateSecretInput{
			Name:         &secretName,
			SecretString: &stringEnvFile,
		})
		if err != nil {
			return fmt.Errorf("creating secret %v", err)
		}
	} else {
		log.Printf("Updating secret %s", secretName)
		_, err = client.PutSecretValue(context.TODO(), &secretsmanager.PutSecretValueInput{
			SecretId:     &secretName,
			SecretString: &stringEnvFile,
		})
		if err != nil {
			return fmt.Errorf("updating secret %v", err)
		}
	}

	return nil

}
