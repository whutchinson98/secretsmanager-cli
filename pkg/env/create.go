package env

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/spf13/cobra"
	"github.com/whutchinson98/secretsmanager-cli/internal/util"
)

var osGetWd = os.Getwd
var osCreate = os.Create

func CreateEnvFile(cmd *cobra.Command, args []string) error {
	region, _ := cmd.Flags().GetString("region")

	cfg, _ := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))

	client := secretsmanager.NewFromConfig(cfg)

	secretName, _ := cmd.Flags().GetString("secret-name")

	fmt.Printf("Secret name %s\n", secretName)

	secVal, err := client.GetSecretValue(context.TODO(), &secretsmanager.GetSecretValueInput{
		SecretId: &secretName,
	})

	if err != nil {
		return fmt.Errorf("fetching secret value: %v", err)
	}

	secretString := aws.ToString(secVal.SecretString)

	sec := map[string]interface{}{}

	if err := json.Unmarshal([]byte(secretString), &sec); err != nil {
		return fmt.Errorf("converting secret value to JSON: %v", err)
	}

	envString := util.BuildEnvStringFromSecret(sec)

	path, _ := cmd.Flags().GetString("path")
	envFileName, _ := cmd.Flags().GetString("env-file")

	envFile, err := InitEnvFile(path, envFileName)

	if err != nil {
		return fmt.Errorf("creating the env file with path %v file name %v: Error %v", path, envFileName, err)
	}

	_, err = envFile.WriteString(envString)

	if err != nil {
		return fmt.Errorf("writing to env file: %v", err)
	}

	return nil
}

func InitEnvFile(path string, envFileName string) (*os.File, error) {

	wd, err := osGetWd()

	if err != nil {
		return nil, err
	}

	if !bytes.HasSuffix([]byte(path), []byte("/")) {
		path = path + "/"
	}

	envFile, err := osCreate(wd + "/" + path + envFileName)

	return envFile, err
}
