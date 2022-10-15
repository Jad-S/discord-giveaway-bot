package config

import (
	sh "giveaway/internal/sheets"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

type EnvironmentVariables struct {
	SheetsCredentials sh.List `yaml:"sheets-credentials"`
	DiscordPublicKey  string  `yaml:"discord-publickey"`
	S3BucketName      string  `yaml:"s3-bucketname"`
}

func LoadAWSEnvironmentVars() EnvironmentVariables {
	return EnvironmentVariables{
		SheetsCredentials: sh.List{
			Guess: sh.Spreadsheet{
				Id:    os.Getenv("GUESS_ID"),
				Range: os.Getenv("GUESS_RANGE"),
			},
			Raffle: sh.Spreadsheet{
				Id:    os.Getenv("RAFFLE_ID"),
				Range: os.Getenv("RAFFLE_RANGE"),
			},
		},
		DiscordPublicKey: os.Getenv("DISCORD_PUBLIC_KEY"),
		S3BucketName:     os.Getenv("S3_BUCKET_NAME"),
	}
}

func LoadLocalEnvironmentVars(filename string) (EnvironmentVariables, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("Error reading file: %s", err)
	}

	configs := EnvironmentVariables{}

	err = yaml.Unmarshal(data, &configs)
	if err != nil {
		log.Printf("in file %q: %v", filename, err)
	}

	return configs, nil
}

func LoadConfigs(file string) (EnvironmentVariables, error) {
	_, exists := os.LookupEnv("AWS_EXECUTION_ENV")
	if exists {
		log.Println("Loading AWS environment variables..")
		return LoadAWSEnvironmentVars(), nil
	}

	log.Println("Loading local environment variables..")

	return LoadLocalEnvironmentVars(file)
}
