package services

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"service_fraud/utils"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

// AwsSecrets stores API keys retrieved from AWS Secrets Manager.
type AwsSecrets struct {
	IpapiKey    string `json:"ipapi_key"`
	CurrencyKey string `json:"currency_key"`
}

var instance *AwsSecrets
var svc *secretsmanager.Client
var onceCreation sync.Once
var onceLoading sync.Once

// NewAwsSecrets returns the singleton instance of AwsSecrets.
func NewAwsSecrets() *AwsSecrets {
	onceCreation.Do(func() {
		instance = &AwsSecrets{}
	})
	return instance
}

// init initializes the AWS Secrets Manager client.
func init() {
	region := "us-east-1"

	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Fatal(err)
	}

	svc = secretsmanager.NewFromConfig(config)
}

// GetSecret retrieves a secret value by name from AWS Secrets Manager.
func (a *AwsSecrets) GetSecret(name string) (*string, error) {
	var loadErr error
	onceLoading.Do(func() {
		input := &secretsmanager.GetSecretValueInput{
			SecretId:     aws.String(utils.SECRET_VAULT),
			VersionStage: aws.String("AWSCURRENT"),
		}

		result, err := svc.GetSecretValue(context.TODO(), input)
		if err != nil {
			log.Println(err.Error())
			loadErr = err
			return
		}

		err = json.Unmarshal([]byte(*result.SecretString), a)
		if err != nil {
			log.Printf("Error unmarshaling JSON: %v", err)
			loadErr = err
			return
		}
		log.Print("Get process secrets successful")
	})

	if loadErr != nil {
		return nil, loadErr
	}

	str := ""
	switch {
	case name == utils.SECRET_API_IP_KEY:
		str = a.IpapiKey
	case name == utils.SECRET_API_CURRENCY_KEY:
		str = a.CurrencyKey
	default:
		log.Printf("The requested value is not valid: %s", name)
		return nil, errors.New(utils.ERR_MESSAGE_GET_SECRETS)
	}

	if str == "" {
		return nil, errors.New(utils.ERR_MESSAGE_GET_SECRETS)
	}
	return &str, nil
}
