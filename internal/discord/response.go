package discord

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

// Response is the response we send back to Discord
type Response struct {
	Type int         `json:"type"`
	Data MessageData `json:"data"`
}

// MessageData contains the message we send back to the Discord channel
type MessageData struct {
	Content string `json:"content"`
}

// Pong This is a PONG response used to acknowledge that the command was received. 
// Check out https://discord.com/developers/docs/interactions/slash-commands#interaction-response-interactioncallbacktype
func Pong() events.APIGatewayProxyResponse {
	var response Response

	response = Response{
		Type: 1,
	}
	fmt.Printf("Response to send: %v", response)

	result, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(result),
	}
}

// Success is the Response we send back to Discord if the request was processed successfully
func Success(b Body) events.APIGatewayProxyResponse {
	var response Response

	response = Response{
		Type: 4,
		Data: MessageData{
			Content: fmt.Sprintf("%s joined the giveaway!", b.Member.User.Name),
		},
	}
	fmt.Printf("Response to send: %v", response)

	result, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(result),
	}
}

func Error(msg string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: 400,
		Body:       msg,
	}
}
