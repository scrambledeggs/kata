// generate-secrets generates a file .secrets.json
// derived from AWS Secrets Manager

package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

const DEFAULT_ENV = "test"
const DEFAULT_REGION = "ap-southeast-1"

func main() {
	projectName := "{{ cookiecutter.project_name }}"
	environment := flag.String("env", DEFAULT_ENV, "Environment")
	region := flag.String("region", DEFAULT_REGION, "AWS Region")
	flag.Parse()

	secretName := *environment + "-" + projectName

	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(*region))
	if err != nil {
		panic(fmt.Sprintf("Error loading config: %v", err))
	}

	svc := secretsmanager.NewFromConfig(config)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"),
	}

	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		panic(fmt.Sprintf("Error getting secret value: %v", err))
	}

	var data map[string]interface{}
	err = json.Unmarshal([]byte(*result.SecretString), &data)
	if err != nil {
		panic(fmt.Sprintf("Error unmarshalling data: %v", err))
	}

	file, _ := json.MarshalIndent(data, "", "  ")
	_ = os.WriteFile(".secrets.json", file, 0644)

	fmt.Printf("Successfully generated .secrets.json from \"%v\"", secretName)
}
