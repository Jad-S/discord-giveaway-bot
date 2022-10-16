package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	cfg "giveaway/config"
	"giveaway/internal/bucket"
	disc "giveaway/internal/discord"
	sh "giveaway/internal/sheets"
	"log"
)

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	configs, err := cfg.LoadConfigs("config.yaml")
	if err != nil {
		log.Printf("Error loading configs: %v", err)
	}
	fmt.Printf("Configs: %+v\n", configs)

	// verify the request is coming from Discord
	verified, _ := disc.VerifyDiscordRequest(request, configs.DiscordPublicKey)
	if !verified {
		return disc.Error("request could not be verified"), err
	}

	// parse the request body
	body, err := disc.ParseRequestBody(request.Body)
	if err != nil {
		return disc.Error("could not parse request body"), err
	}

	// check if this was a ping request from Discord
	if body.Type == 1 {
		return disc.Pong(), nil
	}

	// get the command from the request body
	command := body.Data.CommandName
	fmt.Printf("Processing request for command: %s\n", command)

	var spreadsheet sh.Spreadsheet // spreadsheet to be updated

	// switch on the command to determine which spreadsheet to use
	switch command {
	case "raffle":
		spreadsheet = configs.SheetsCredentials.Raffle
	case "guess":
		spreadsheet = configs.SheetsCredentials.Guess
	}

	var file []byte

	// pull service credentials file for our Google service account from S3
	file, err = bucket.GetFile("client.json")
	if err != nil {
		return disc.Error("error getting credentials file"), err
	}

	// begin processing the request in sheets
	client := sh.NewClient(file)
	log.Println("Created client, updating spreadsheet...")

	// check if the user is already in the spreadsheet
	userExists := sh.VerifyEntry(client, spreadsheet, body.Member.User.Id)
	if userExists {
		return disc.Error(fmt.Sprintf("error: Duplicate entry, nice try %s", body.Member.User.Name)), err
	}

	// update the spreadsheet
	err = sh.AddNewEntry(client, body, spreadsheet)
	if err != nil {
		return disc.Error("error adding new entry"), err
	}

	return disc.Success(body), nil
}
